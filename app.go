package main

import (
	"html/template"
	"log"
	"masterjulz/lucullus/api"
	Honey "masterjulz/lucullus/honeypotconn"
	OW "masterjulz/lucullus/openwhisk"
	WIOT "masterjulz/lucullus/wiot"
	"net/http"
	"os"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
	index.Execute(w, nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	api.RegisterHandlers()

	http.HandleFunc("/", helloworld)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	go http.ListenAndServe(":"+port, nil)

	ow := OW.New(
		"MjAxNTU1YjMtMDlhYy00MDI1LTk3OGMtMjJhYWZiYjM4OThiOnVRbTFMQmdPNzVtMWl5Q1prVE10S3NPQTBOWWhmdGxWMTJsY1ZFd29KRVpnMEgyWFRYbG95cjlEWEZNRlh5emg=",
		"skupnjak%40de.ibm.com_SCM-Project",
	)

	honey := Honey.New(ow)

	wiot := WIOT.New(
		"tcp://y1iv4w.messaging.internetofthings.ibmcloud.com:1883",
		"iot-2/type/Sensor/id/lucullus/evt/input/fmt/json",
		"a:y1iv4w:lucullus",
		"a-y1iv4w-pogrhsaynn",
		"s*ay_o6rkeH+W7Exsm",
	)

	wiot.Connect()
	go wiot.Subscribe(honey)
	//wiot.Publish5Messages()
	go honey.Subscribe("lucullus")

	select {}

	wiot.Unsubscribe()
	wiot.Disconnect()
}
