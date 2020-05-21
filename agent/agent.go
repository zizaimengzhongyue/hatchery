package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zizaimengzhongyue/serverr-manager/types"
)

func Register(name, host string, port int) {
	service := types.Service{
		Name: name,
		Host: host,
		Port: port,
	}
	client := &http.Client{}
	body, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8001", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Body = bytes.NewReader(body)
	_, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}
