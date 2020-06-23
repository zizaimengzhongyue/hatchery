package agent

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zizaimengzhongyue/hatchery/types"
)

var cfg types.Config = &types.Config{}
var servers map[string][]types.Server = map[string][]types.Server{}

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

func Sync(names []string) error {
	params = strings.join(names, ",")
	url := "http://127.0.0.1:8001/get?names=" + params
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	err := json.Unmarshal(body, servers)
	if err != nil {
		return nil
	}
}

func Init(names []string) error {
}

func InitWithFile(path string) error {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bts, cfg)
	if err != nil {
		return err
	}
    ticker := time.NewTicker(60 * time.Second)
	go func() {
		for {
			select {
            case t := <-ticker.C:
				if err = Sync(cfg.names); err != nil {
					return err
				}
			}
		}
	}()
	return nil
}

func DoRequest(name string, data interface{}, res interface{}) error {
	return nil
}
