package api

import (
	"fmt"
	WIOT "masterjulz/lucullus/wiot"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var apiPrefix = "/api/v1/"
var filterURL = "http://riddle-filter.ng.mybluemix.net"

type Filter struct {
	Active              bool    `json:"active"`
	Topic               string  `json:"topic"`
	FilterTopId         string  `json:"filter-top-id"`
	FilterBottomId      string  `json:"filter-bottom-id"`
	TresholdValueTop    float64 `json:"treshold-value-top"`
	TresholdValueButtom float64 `json:"treshhold-value-bottom"`
}

func RegisterHandlers(wiot *WIOT.WatsonIoT) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(apiPrefix+"start", startHandlerChain).Methods("POST")
	router.HandleFunc(apiPrefix+"stop", stopHandlerChain).Methods("POST")
	router.HandleFunc(apiPrefix+"publish", makeWiotHandler(publish, wiot)).Methods("POST")
	http.Handle(apiPrefix, router)
	handler := cors.Default().Handler(router)
	return handler
}

func makeWiotHandler(fn func(http.ResponseWriter, *http.Request, *WIOT.WatsonIoT), wiot *WIOT.WatsonIoT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, wiot)
	}
}

func startHandlerChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting handler chain... ")
}

func stopHandlerChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stopping handler chain... ")
}

func publish(w http.ResponseWriter, r *http.Request, wiot *WIOT.WatsonIoT) {
	wiot.Publish5Messages()
}
