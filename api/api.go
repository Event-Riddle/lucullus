package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var apiPrefix = "/api/v1/"
var filterURL = "http://riddle-filter.ng.mybluemix.net"

func RegisterHandlers() {
	router := mux.NewRouter()
	router.HandleFunc(apiPrefix+"start", startHandlerChain).Methods("POST")
	http.Handle(apiPrefix, router)
}

func startHandlerChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it WORX DUDE")
}
