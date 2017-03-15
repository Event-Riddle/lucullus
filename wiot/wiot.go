package wiot

import (
	"fmt"
	Honey "masterjulz/lucullus/honeypotconn"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type WatsonIoT struct {
	Client MQTT.Client
	Topic  string
}

func getMqttOpts(url string, clientID string, username string, password string) *MQTT.ClientOptions {
	fmt.Printf("url: %s, clientID: %s, username: %s, password: %s", url, clientID, username, password)
	opts := MQTT.NewClientOptions().AddBroker(url)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetCleanSession(true) //Check
	return opts
}

func New(url string, topic string, clientID string, username string, password string) *WatsonIoT {
	opts := getMqttOpts(url, clientID, username, password)
	wiot := WatsonIoT{
		MQTT.NewClient(opts),
		topic,
	}
	return &wiot
}

func (wiot *WatsonIoT) Connect() {
	if token := wiot.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("App Client connected")
}

func (wiot *WatsonIoT) Subscribe(honey *Honey.HoneyPot) {
	if token := wiot.Client.Subscribe(wiot.Topic, 1, HPotMessageHandler(sayHello, honey)); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		fmt.Println("subscription error")
		os.Exit(1)
	}
}

func (wiot *WatsonIoT) Unsubscribe() {
	if token := wiot.Client.Unsubscribe(wiot.Topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("Unsubscribed")
}

func (wiot *WatsonIoT) Disconnect() {
	wiot.Client.Disconnect(250)
	fmt.Println("disconnected")
}

func HPotMessageHandler(fn func(MQTT.Client, MQTT.Message, *Honey.HoneyPot), honey *Honey.HoneyPot) MQTT.MessageHandler {
	return func(client MQTT.Client, msg MQTT.Message) {
		fn(client, msg, honey)
	}
}

func sayHello(client MQTT.Client, msg MQTT.Message, honey *Honey.HoneyPot) {
	payload := string(msg.Payload())
	fmt.Printf("the message is %s\n", payload)
	fmt.Printf("the messsage is %s", msg.Payload())
	honey.Publish("filter", payload)
}

func (wiot *WatsonIoT) Publish5Messages() {
	for i := 0; i < 5; i++ {
		//	text  := fmt.Println("{\"msg\":\"hallo\"}")
		token := wiot.Client.Publish(wiot.Topic, 1, false, "{\"Degree\":30}")

		token.Wait()
	}
	fmt.Println("Published 5 messages")
}
