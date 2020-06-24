package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zizaimengzhongyue/hatchery/types"
)

var cfg types.Config = types.Config{}
var remoterMultiResponse types.RemoterMultiResponse = types.RemoterMultiResponse{}
var servers map[string][]types.Service = map[string][]types.Service{}

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
	remoterRes := &types.RemoterMultiResponse{}
	params := strings.Join(names, ",")
	url := "http://127.0.0.1:8001/multiGet?names=" + params
	log.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	err = json.Unmarshal(body, remoterRes)
	if err != nil {
		return err
	}
	servers = remoterRes.Data
	return nil
}

func Init(names []string) error {
	return nil
}

func InitWithFile(path string) error {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bts, &cfg.Names)
	if err != nil {
		return err
	}
	err = Sync(cfg.Names)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(servers)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			_ = <-ticker.C
			err = Sync(cfg.Names)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(servers)
		}
	}()
	return nil
}

func DoRequest(name string, data interface{}, res interface{}) error {
	return nil
}
