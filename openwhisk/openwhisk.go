package openwhisk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenWhisk struct {
	BaseURL   string
	Token     string
	Namespace string

	Client http.Client
}

func New(token string, namespace string) *OpenWhisk {
	ow := OpenWhisk{
		"https://openwhisk.ng.bluemix.net/api/v1",
		token,
		namespace,
		http.Client{},
	}
	return &ow
}

func (ow *OpenWhisk) TriggerAction(action string, args string) {
	url := ow.BaseURL + "/namespaces/" + ow.Namespace + "/actions/" + action
	jsonStr := []byte(args)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	addHeader(req, ow.Token)
	client := ow.Client
	doRequest(req, client)
}

func doRequest(req *http.Request, client http.Client) {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing POST request")
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("OPEN WHIS RESPONSE: %s", body)
}

func addHeader(req *http.Request, token string) {
	auth := "Basic " + token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)
}
