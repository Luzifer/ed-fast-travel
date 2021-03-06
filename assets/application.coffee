searchTimer = null
systemCache = null
route = []
sock = null
firstConnect = true
  
$ ->
  bindUI()
  loadSettings()
  spawnSocket()

bindUI = () ->
  $('#settings-save').bind 'click', saveSettings
  $('#inputSystem').bind 'input', handleSystemInputEvent
  $('#inputSystem').bind 'keyup', handleSystemKeyUpEvent
  $('#add-system-btn').bind 'click', handleSystemAdd
  $('#update-system').bind 'click', handleUpdate
  $('#update-database').bind 'click', handleUpdateDatabase
  $('#shutdown-system').bind 'click', handleShutdown
  $(window).bind 'hashchange', handleHashChange

loadSettings = () ->
  settings = localStorage.getItem 'settings'
  s =
    stop_range: 500
    show_coordinates: true

  if settings != null
    s = JSON.parse settings

  $('#stop-range').val s.stop_range
  $('#show-coords').prop 'checked', s.show_coordinates
  if !s.show_coordinates
    $('body').addClass 'no-coords'
  else
    $('body').removeClass 'no-coords'

saveSettings = () ->
  s =
    stop_range: $('#stop-range').val()
    show_coordinates: $('#show-coords').prop 'checked'

  localStorage.setItem 'settings', JSON.stringify(s)
  $('#settings').modal 'hide'

  if !s.show_coordinates
    $('body').addClass 'no-coords'
  else
    $('body').removeClass 'no-coords'

setInputState = (state) ->
  iconClass = switch state
    when 'searching' then 'fa fa-spinner fa-pulse'
    when 'error' then 'fa fa-exclamation-triangle'
    when 'success' then 'fa fa-check'
    else 'fa fa-pencil'
  $('#add-system-icon').attr 'class', iconClass

setWarning = (warning) ->
  if warning == ""
    $('#warning').hide()
  else
    $('#warning-text').text warning
    $('#warning').show()

handleSystemInputEvent = () ->
  setInputState 'writing'
  $('#add-system-btn').prop 'disabled', true

  if searchTimer != null
    window.clearTimeout searchTimer
    searchTimer = null

  if $('#inputSystem').val() != ''
    searchTimer = window.setTimeout searchSystemByName, 1500

handleSystemKeyUpEvent = (evt) ->
  if evt.keyCode == 13
    $('#add-system-btn').trigger 'click'

searchSystemByName = () ->
  setInputState 'searching'
  searchString = $('#inputSystem').val()

  $.get "/api/system-by-name?system_name=#{searchString}", (data) ->
    setWarning data.error_message
    if !data.success
      setInputState 'error'
      $('#add-system-group').addClass('has-error')
    else
      setInputState 'success'
      $('#add-system-group').removeClass('has-error')
      $('#inputSystem').val data.data.system.name
      $('#add-system-btn').prop 'disabled', false
      systemCache = data.data.system

handleSystemAdd = () ->
  if $('#add-system-btn').prop 'disabled' or systemCache == null
    return

  if route[0] == '' and route.length == 1
    route = []

  if route[route.length - 1] != systemCache.id
    route.push systemCache.id
    location.hash = route.join(',')

  $('#inputSystem').val ''
  $('#inputSystem').trigger 'input'
  systemCache = null

spawnSocket = () ->
  endpoint = switch location.protocol
    when "https:" then "wss://#{location.host}/api/route"
    else "ws://#{location.host}/api/route"

  sock = new WebSocket(endpoint)
  sock.onclose = ->
    setWarning window.translations["warn_communication_lost"]
  sock.onmessage = handleRouteResult
  sock.onopen = () ->
    if firstConnect
      $(window).trigger 'hashchange'
      firstConnect = false

handleRouteResult = (evt) ->
  data = JSON.parse(evt.data)

  if !data.success
    setWarning data.error_message
    return

  pbar = null

  for el in $('tr.auto-added.pbar')
    if $(el).data('request_id') != "#{data.route_request_id}"
      continue
    pbar = el

  if data.result.progress < 1
    $(pbar).find('.progress-bar').css('width', "#{data.result.progress * 100}%")

  if data.result.trace_type != "flight_stop"
    return

  editRow = null

  for el in $('tr.auto-added')
    if $(el).data('request_id') != "#{data.route_request_id}"
      continue
    if $(el).data('system_id') != "#{data.result.star_system.id}"
      continue
    editRow = el
    break

  if editRow == null
    editRow = getSystemLine()
    $(editRow).insertBefore(pbar)

  if data.counter > 0
    $(editRow).find('.stop_no').text(data.counter)
  $(editRow).find('.system_name').text(data.result.star_system.name)
  $(editRow).find('.coordinates').text(compileCoordinates(data.result.star_system.coords))
  $(editRow).find('.flight_distance').text("#{data.result.flight_distance.toFixed(2)} Ly")
  $(editRow).find('.total_flight_distance').text("#{data.result.total_flight_distance.toFixed(2)} Ly")
  $(editRow).find('.clip').attr 'data-clipboard-text', data.result.star_system.name
  if !data.result.star_system.scoopable
    $(editRow).find('.info-scoop').remove()
  if !data.result.star_system.permit
    $(editRow).find('.info-permit').remove()
  clip = new Clipboard $(editRow).find('.clip')[0]
  clip.on 'success', (e) ->
    toggleButtonClass e.trigger, 'btn-default', 'btn-success'
    delay 1500, () ->
      toggleButtonClass e.trigger, 'btn-success', 'btn-default'

  if data.result.progress == 1
    $(pbar).remove()

compileCoordinates = (coords) ->
  "(#{coords.x.toFixed(5)} / #{coords.y.toFixed(5)} / #{coords.z.toFixed(5)})"

handleHashChange = () ->
  $('.auto-added').remove()
  route = location.hash.substring(1).split(',')

  if location.hash.substring(1) == ""
    return

  tmp = route
  ct = 0
  while tmp.length > 1
    request_id = "#{tmp[0]}::#{tmp[1]}::#{ct}"

    pbar = getProgressBar()
    $(pbar).data 'request_id', request_id
    $(pbar).insertBefore $('#inputSystemRow')

    msg =
      start_system_id: parseInt(tmp[0])
      target_system_id: parseInt(tmp[1])
      route_request_id: request_id
      stop_distance: parseFloat($('#stop-range').val())
    sock.send(JSON.stringify(msg))
    tmp = tmp[1..]
    ct = ct + 1

getProgressBar = () ->
  $('<tr class="auto-added pbar">
       <td class="right">&nbsp;</td>
       <td>
         <div class="progress">
           <div class="progress-bar progress-bar-striped active" role="progressbar" style="width: 0%"></div>
         </div>
       </td>
       <td class="right">&nbsp;</td>
       <td class="right">&nbsp;</td>
     </tr>')

getSystemLine = () ->
  $('<tr class="auto-added">
       <td class="right stop_no"></td>
       <td>
        <button class="btn btn-default btn-xs clip"><i class="fa fa-clipboard" aria-hidden="true"></i></button>
         <i class="fa fa-bolt squareinfo info-scoop" title="Scoopable star"></i>
         <i class="fa fa-ticket squareinfo info-permit" title="Requires permit"></i>
         <span class="system_name"></span>
         <span class="coordinates"></span>
       </td>
       <td class="right flight_distance"></td>
       <td class="right total_flight_distance"></td>
     </tr>')

handleShutdown = () ->
  $.get '/api/control/shutdown', (data) ->
    setWarning data.error_message
  
handleUpdate = () ->
  $.get '/api/control/update', (data) ->
    setWarning data.error_message

handleUpdateDatabase = () ->
  $.get '/api/control/update-database', (data) ->
    setWarning data.error_message

toggleButtonClass = (button, remove, add) ->
  $(button).removeClass remove
  $(button).addClass add

delay = (time, fkt) ->
  window.setTimeout fkt, time
