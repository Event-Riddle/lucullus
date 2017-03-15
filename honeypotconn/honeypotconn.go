package honeypotconn

import (
	"fmt"
	OW "masterjulz/lucullus/openwhisk"
	"time"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
	"github.com/streadway/amqp"
)

type HoneyPot struct {
	URL        string
	Connection *amqp.Connection
	pubChannel *amqp.Channel
	subChannel *amqp.Channel
	openWhisk  *OW.OpenWhisk
}

func New(ow *OW.OpenWhisk) *HoneyPot {
	url := "amqp://vvesrlkq:7cTOIc7-W2awpfANfNqHsFx7tMfocTds@white-swan.rmq.cloudamqp.com/vvesrlkq"
	conn, _ := amqp.Dial(url)
	sub, _ := conn.Channel()
	pub, _ := conn.Channel()
	hp := HoneyPot{
		url,
		conn,
		pub,
		sub,
		ow,
	}
	return &hp
}

func (hp *HoneyPot) Subscribe(topic string) {
	channel := hp.subChannel
	autoAck, exclusive, noLocal, noWait := false, false, false, false
	messages, _ := channel.Consume(topic, "#", autoAck, exclusive, noLocal, noWait, nil)
	multiAck := false
	in := 1
	for msg := range messages {
		in++
		fmt.Println("Body:", string(msg.Body), "Timestamp:", msg.Timestamp, in)
		msg.Ack(multiAck)
		hp.openWhisk.TriggerAction("hello", `{"arg":"moto"}`)
	}
}

func (hp *HoneyPot) Publish(key string, payload string) {
	// for t := range timer.C {
	msg := amqp.Publishing{
		DeliveryMode: 1,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         []byte(payload),
	}
	mandatory, immediate := false, false
	hp.pubChannel.Publish("amq.topic", key, mandatory, immediate, msg)
	// }
}

func (hp *HoneyPot) CloseConnection() {
	hp.Connection.Close()
	fmt.Println("amqp connection closed")
}

func (hp *HoneyPot) ClosePubChannel() {
	hp.pubChannel.Close()
	fmt.Println("amqp publish channel closed")
}

func (hp *HoneyPot) CloseSubChannel() {
	hp.subChannel.Close()
	fmt.Println("amqp subscribtion channel closed")
}

func (hp *HoneyPot) CloseAll() {
	hp.ClosePubChannel()
	hp.CloseSubChannel()
	hp.CloseConnection()
}
