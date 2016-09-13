// Generated by CoffeeScript 1.10.0
(function() {
  var bindUI, compileCoordinates, delay, firstConnect, getProgressBar, getSystemLine, handleHashChange, handleRouteResult, handleShutdown, handleSystemAdd, handleSystemInputEvent, handleUpdate, loadSettings, route, saveSettings, searchSystemByName, searchTimer, setInputState, setWarning, sock, spawnSocket, systemCache, toggleButtonClass;

  searchTimer = null;

  systemCache = null;

  route = [];

  sock = null;

  firstConnect = true;

  $(function() {
    bindUI();
    loadSettings();
    return spawnSocket();
  });

  bindUI = function() {
    $('#settings-save').bind('click', saveSettings);
    $('#inputSystem').bind('input', handleSystemInputEvent);
    $('#add-system-btn').bind('click', handleSystemAdd);
    $('#update-system').bind('click', handleUpdate);
    $('#shutdown-system').bind('click', handleShutdown);
    return $(window).bind('hashchange', handleHashChange);
  };

  loadSettings = function() {
    var s, settings;
    settings = localStorage.getItem('settings');
    s = {
      stop_range: 500,
      show_coordinates: true
    };
    if (settings !== null) {
      s = JSON.parse(settings);
    }
    $('#stop-range').val(s.stop_range);
    $('#show-coords').prop('checked', s.show_coordinates);
    if (!s.show_coordinates) {
      return $('body').addClass('no-coords');
    } else {
      return $('body').removeClass('no-coords');
    }
  };

  saveSettings = function() {
    var s;
    s = {
      stop_range: $('#stop-range').val(),
      show_coordinates: $('#show-coords').prop('checked')
    };
    localStorage.setItem('settings', JSON.stringify(s));
    $('#settings').modal('hide');
    if (!s.show_coordinates) {
      return $('body').addClass('no-coords');
    } else {
      return $('body').removeClass('no-coords');
    }
  };

  setInputState = function(state) {
    var iconClass;
    iconClass = (function() {
      switch (state) {
        case 'searching':
          return 'fa fa-spinner fa-pulse';
        case 'error':
          return 'fa fa-exclamation-triangle';
        case 'success':
          return 'fa fa-check';
        default:
          return 'fa fa-pencil';
      }
    })();
    return $('#add-system-icon').attr('class', iconClass);
  };

  setWarning = function(warning) {
    if (warning === "") {
      return $('#warning').hide();
    } else {
      $('#warning-text').text(warning);
      return $('#warning').show();
    }
  };

  handleSystemInputEvent = function() {
    setInputState('writing');
    $('#add-system-btn').prop('disabled', true);
    if (searchTimer !== null) {
      window.clearTimeout(searchTimer);
      searchTimer = null;
    }
    if ($('#inputSystem').val() !== '') {
      return searchTimer = window.setTimeout(searchSystemByName, 1500);
    }
  };

  searchSystemByName = function() {
    var searchString;
    setInputState('searching');
    searchString = $('#inputSystem').val();
    return $.get("/api/system-by-name?system_name=" + searchString, function(data) {
      setWarning(data.error_message);
      if (!data.success) {
        setInputState('error');
        return $('#add-system-group').addClass('has-error');
      } else {
        setInputState('success');
        $('#add-system-group').removeClass('has-error');
        $('#inputSystem').val(data.data.system.name);
        $('#add-system-btn').prop('disabled', false);
        return systemCache = data.data.system;
      }
    });
  };

  handleSystemAdd = function() {
    if ($('#add-system-btn').prop('disabled' || systemCache === null)) {
      return;
    }
    if (route[route.length - 1] !== systemCache.id) {
      route.push(systemCache.id);
      location.hash = route.join(',');
    }
    $('#inputSystem').val('');
    $('#inputSystem').trigger('input');
    return systemCache = null;
  };

  spawnSocket = function() {
    var endpoint;
    endpoint = (function() {
      switch (location.protocol) {
        case "https:":
          return "wss://" + location.host + "/api/route";
        default:
          return "ws://" + location.host + "/api/route";
      }
    })();
    sock = new WebSocket(endpoint);
    sock.onclose = function() {
      return setWarning('Communication with server ended unexpectedly. Please reload this page to continue your calculation.');
    };
    sock.onmessage = handleRouteResult;
    return sock.onopen = function() {
      if (firstConnect) {
        $(window).trigger('hashchange');
        return firstConnect = false;
      }
    };
  };

  handleRouteResult = function(evt) {
    var clip, data, editRow, el, i, j, len, len1, pbar, ref, ref1;
    data = JSON.parse(evt.data);
    if (!data.success) {
      setWarning(data.error_message);
      return;
    }
    pbar = null;
    ref = $('tr.auto-added.pbar');
    for (i = 0, len = ref.length; i < len; i++) {
      el = ref[i];
      if ($(el).data('request_id') !== ("" + data.route_request_id)) {
        continue;
      }
      pbar = el;
    }
    editRow = null;
    ref1 = $('tr.auto-added');
    for (j = 0, len1 = ref1.length; j < len1; j++) {
      el = ref1[j];
      if ($(el).data('request_id') !== ("" + data.route_request_id)) {
        continue;
      }
      if ($(el).data('system_id') !== ("" + data.result.star_system.id)) {
        continue;
      }
      editRow = el;
      break;
    }
    if (editRow === null) {
      editRow = getSystemLine();
      $(editRow).insertBefore(pbar);
    }
    if (data.counter > 0) {
      $(editRow).find('.stop_no').text(data.counter);
    }
    $(editRow).find('.system_name').text(data.result.star_system.name);
    $(editRow).find('.coordinates').text(compileCoordinates(data.result.star_system.coords));
    $(editRow).find('.flight_distance').text((data.result.flight_distance.toFixed(2)) + " Ly");
    $(editRow).find('.total_flight_distance').text((data.result.total_flight_distance.toFixed(2)) + " Ly");
    $(editRow).find('.clip').attr('data-clipboard-text', data.result.star_system.name);
    clip = new Clipboard($(editRow).find('.clip')[0]);
    clip.on('success', function(e) {
      toggleButtonClass(e.trigger, 'btn-default', 'btn-success');
      return delay(1500, function() {
        return toggleButtonClass(e.trigger, 'btn-success', 'btn-default');
      });
    });
    if (data.result.progress < 1) {
      return $(pbar).find('.progress-bar').css('width', (data.result.progress * 100) + "%");
    } else {
      return $(pbar).remove();
    }
  };

  compileCoordinates = function(coords) {
    return "(" + (coords.x.toFixed(5)) + " / " + (coords.y.toFixed(5)) + " / " + (coords.z.toFixed(5)) + ")";
  };

  handleHashChange = function() {
    var ct, msg, pbar, request_id, results, start, target, tmp;
    if (location.hash.substring(1) === "") {
      return;
    }
    $('.auto-added').remove();
    route = location.hash.substring(1).split(',');
    tmp = route;
    ct = 0;
    results = [];
    while (tmp.length > 1) {
      request_id = tmp[0] + "::" + tmp[1] + "::" + ct;
      start = getSystemLine();
      $(start).data('request_id', request_id);
      $(start).data('system_id', tmp[0]);
      $(start).insertBefore($('#inputSystemRow'));
      pbar = getProgressBar();
      $(pbar).data('request_id', request_id);
      $(pbar).insertBefore($('#inputSystemRow'));
      target = getSystemLine();
      $(target).data('request_id', request_id);
      $(target).data('system_id', tmp[1]);
      $(target).insertBefore($('#inputSystemRow'));
      msg = {
        start_system_id: parseInt(tmp[0]),
        target_system_id: parseInt(tmp[1]),
        route_request_id: request_id,
        stop_distance: parseFloat($('#stop-range').val())
      };
      sock.send(JSON.stringify(msg));
      tmp = tmp.slice(1);
      results.push(ct = ct + 1);
    }
    return results;
  };

  getProgressBar = function() {
    return $('<tr class="auto-added pbar"> <td class="right">&nbsp;</td> <td> <div class="progress"> <div class="progress-bar progress-bar-striped active" role="progressbar" style="width: 0%"></div> </div> </td> <td class="right">&nbsp;</td> <td class="right">&nbsp;</td> </tr>');
  };

  getSystemLine = function() {
    return $('<tr class="auto-added"> <td class="right stop_no"></td> <td> <button class="btn btn-default btn-xs clip"><i class="fa fa-clipboard" aria-hidden="true"></i></button> <span class="system_name"></span> <span class="coordinates"></span> </td> <td class="right flight_distance"></td> <td class="right total_flight_distance"></td> </tr>');
  };

  handleShutdown = function() {
    return $.get('/api/control/shutdown', function(data) {
      return setWarning(data.error_message);
    });
  };

  handleUpdate = function() {
    return $.get('/api/control/update', function(data) {
      return setWarning(data.error_message);
    });
  };

  toggleButtonClass = function(button, remove, add) {
    $(button).removeClass(remove);
    return $(button).addClass(add);
  };

  delay = function(time, fkt) {
    return window.setTimeout(fkt, time);
  };

}).call(this);
