package agent

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8001/register", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return nil
}
