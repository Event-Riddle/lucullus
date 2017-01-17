package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func RegisterHandlers() {
	router := mux.NewRouter()
	router.HandleFunc(apiPrefix+"start", startHandlerChain).Methods("POST")
	http.Handle(apiPrefix, router)
}

func startHandlerChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it WORX DUDE")
}
