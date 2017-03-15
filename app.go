package main

import (
	"fmt"
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

	wiot := WIOT.New(
		"tcp://y1iv4w.messaging.internetofthings.ibmcloud.com:1883",
		"iot-2/type/Sensor/id/lucullus/evt/input/fmt/json",
		"a:y1iv4w:lucullus",
		"a-y1iv4w-pogrhsaynn",
		"s*ay_o6rkeH+W7Exsm",
	)

	handler := api.RegisterHandlers(wiot)

	http.HandleFunc("/", helloworld)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	go http.ListenAndServe(":"+port, handler)

	ow := OW.New(
		"MDA1YTYwZDAtYTNjYS00YjRlLTljM2ItZDU4ZjY3Nzc3ZDVkOlBtelRkbWRoVFpzWXYwN1dFT3FZMU83Q0FOZXRZY3Nqd0hLVVhiSmNIVXJ4RTVvYVhvUmRhWHJRVDlMSFBMQXI=",
		"skupnjak%40de.ibm.com_SCM-Project",
	)

	honey := Honey.New(ow)

	wiot.Connect()
	fmt.Println("Subscribe to Watson IoT")
	go wiot.Subscribe(honey)
	// wiot.Publish5Messages()
	fmt.Println("subscribe to lucullus topic")
	go honey.Subscribe("lucullus")
	fmt.Println("Subscribed...")
	// wiot.Publish5Messages()
	select {}

	wiot.Unsubscribe()
	wiot.Disconnect()
}
