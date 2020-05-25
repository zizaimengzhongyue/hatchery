package agent

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/zizaimengzhongyue/hatchery/types"
)

func Register(name, host string, port int) error {
	service := types.Service{
		Name: name,
		Host: host,
		Port: port,
	}
	client := &http.Client{}
	body, _ := json.Marshal(service)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8001", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	_, err := client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
