package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Luzifer/gobuilder/autoupdate"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nicksnyder/go-i18n/i18n"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Number of concurrent calculations per socket
	calculationPoolSize = 3
)

var upgrader = websocket.Upgrader{}

type routeRequest struct {
	StartSystemID  int64   `json:"start_system_id"`
	TargetSystemID int64   `json:"target_system_id"`
	RouteRequestID string  `json:"route_request_id"`
	StopDistance   float64 `json:"stop_distance"`
}

type routeResponse struct {
	Counter        int64       `json:"counter"`
	Success        bool        `json:"success"`
	ErrorMessage   string      `json:"error_message"`
	RouteRequestID string      `json:"route_request_id"`
	Result         traceResult `json:"result"`
}

type jsonResponse struct {
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"error_message"`
	Data         map[string]interface{} `json:"data"`
}

func (j jsonResponse) Send(res http.ResponseWriter, cachingAllowed bool) error {
	res.Header().Set("Content-Type", "application/json")
	if cachingAllowed {
		res.Header().Set("Cache-Control", "public, max-age=3600")
	} else {
		res.Header().Set("Cache-Control", "no-cache")
	}

	return json.NewEncoder(res).Encode(j)
}

func startWebService() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleFrontend)
	r.HandleFunc("/assets/application.js", handleJS)

	r.HandleFunc("/api/system-by-name", handleSystemByName)
	r.HandleFunc("/api/route", handleRouteSocket)

	r.HandleFunc("/api/control/shutdown", handleShutdown)
	r.HandleFunc("/api/control/update", handleUpdate)

	log.Fatalf("Unable to listen for web connections: %s", http.ListenAndServe(cfg.Listen, r))
}

func getTranslator(r *http.Request) i18n.TranslateFunc {
	c, _ := r.Cookie("lang")
	var cookieLang string
	if c != nil {
		cookieLang = c.Value
	}
	acceptLang := r.Header.Get("Accept-Language")
	defaultLang := "en-US" // known valid language
	T, _ := i18n.Tfunc(cookieLang, acceptLang, defaultLang)
	return T
}

func handleFrontend(res http.ResponseWriter, r *http.Request) {
	T := getTranslator(r)

	frontend, err := Asset("assets/frontend.html")
	if err != nil {
		http.Error(res, "Could not load frontend: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("frontend").Funcs(template.FuncMap{"T": T}).Parse(string(frontend))
	if err != nil {
		http.Error(res, "Could not parse frontend: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(res, map[string]interface{}{
		"version":                version,
		"disableSoftwareControl": cfg.DisableSoftwareControl,
	}); err != nil {
		http.Error(res, "Could not execute frontend: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleJS(res http.ResponseWriter, r *http.Request) {
	js, _ := Asset("assets/application.js")

	res.Header().Set("Content-Type", "application/javascript")

	res.Write(js)
}

func handleShutdown(res http.ResponseWriter, r *http.Request) {
	T := getTranslator(r)
	if cfg.DisableSoftwareControl {
		http.Error(res, "Controls are disabled", http.StatusForbidden)
		return
	}

	defer os.Exit(0)

	jsonResponse{
		Success:      true,
		ErrorMessage: T("warn_service_will_shutdown"),
	}.Send(res, false)

	<-time.After(time.Second) // Give the response a second to send
}

func handleUpdate(res http.ResponseWriter, r *http.Request) {
	T := getTranslator(r)
	if cfg.DisableSoftwareControl {
		http.Error(res, "Controls are disabled", http.StatusForbidden)
		return
	}

	if hasUpdate, err := autoupdate.New(autoUpdateRepo, autoUpdateLabel).HasUpdate(); err != nil {
		jsonResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}.Send(res, false)
		return
	} else {
		if !hasUpdate {
			jsonResponse{
				Success:      false,
				ErrorMessage: T("warn_no_new_version_found"),
			}.Send(res, false)
			return
		}
	}

	if err := autoupdate.New(autoUpdateRepo, autoUpdateLabel).SingleRun(); err != nil {
		jsonResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}.Send(res, false)
	} else {
		jsonResponse{
			Success:      true,
			ErrorMessage: T("warn_service_update_success"),
		}.Send(res, false)
	}
}

func handleSystemByName(res http.ResponseWriter, r *http.Request) {
	T := getTranslator(r)
	search := r.URL.Query().Get("system_name")

	if len(search) < 3 {
		jsonResponse{
			Success:      false,
			ErrorMessage: T("warn_too_few_characers"),
		}.Send(res, true)
		return
	}

	if sys := starSystems.GetSystemByName(search); sys != nil {
		jsonResponse{
			Success: true,
			Data: map[string]interface{}{
				"system": sys,
			},
		}.Send(res, true)
	} else {
		jsonResponse{
			Success:      false,
			ErrorMessage: T("warn_no_matching_system_found"),
		}.Send(res, true)
		return
	}
}

func handleRouteSocket(res http.ResponseWriter, r *http.Request) {
	T := getTranslator(r)
	// In case socket quits also quit all child operations
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ws, err := upgrader.Upgrade(res, r, nil)
	if err != nil {
		http.Error(res, "Could not open socket", http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	doneChan := make(chan struct{})
	defer close(doneChan)
	go pingSocket(ws, doneChan)

	messageChan := make(chan routeResponse, 500)
	defer close(messageChan)
	go func(ws *websocket.Conn, m chan routeResponse) {
		for msg := range m {
			ws.WriteJSON(msg)
		}
	}(ws, messageChan)

	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	calculationPool := make(chan *struct{}, calculationPoolSize)

	for {
		msg := routeRequest{}

		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("Experienced error: %s", err)
			break
		}

		if msg.RouteRequestID == "" || msg.StartSystemID == 0 || msg.TargetSystemID == 0 {
			messageChan <- routeResponse{
				Success:      false,
				ErrorMessage: T("warn_required_field_missing"),
			}
			continue
		}

		if msg.StopDistance < cfg.WebRouteStopMin {
			messageChan <- routeResponse{
				Success:      false,
				ErrorMessage: T("warn_stop_distance_too_small", cfg),
			}
			continue
		}

		go processSocketRouting(ctx, messageChan, msg, calculationPool)
	}

	cancel()
}

func processSocketRouting(parentCtx context.Context, msgChan chan routeResponse, r routeRequest, calculationPool chan *struct{}) {
	start := starSystems.GetSystemByID(r.StartSystemID)
	target := starSystems.GetSystemByID(r.TargetSystemID)

	calculationPool <- nil
	defer func() { <-calculationPool }()

	ctx, cancel := context.WithTimeout(parentCtx, cfg.WebRouteTimeout)
	defer cancel()

	rChan, eChan := starSystems.CalculateRoute(ctx, start, target, r.StopDistance)

	var counter int64

	for {
		select {
		case stop, ok := <-rChan:
			if ok {
				msgChan <- routeResponse{
					Counter:        counter,
					Success:        true,
					RouteRequestID: r.RouteRequestID,
					Result:         stop,
				}
				counter++
			} else {
				return
			}
		case err := <-eChan:
			if err != nil {
				msgChan <- routeResponse{
					Success:      false,
					ErrorMessage: err.Error(),
				}
				return
			}
		}
	}
}

func pingSocket(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
		case <-done:
			return
		}
	}
}
