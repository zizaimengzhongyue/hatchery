package agent

import (
	"bytes"
	"encoding/json"
    "errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/zizaimengzhongyue/hatchery/types"
)

var cfg types.Config = types.Config{}
var remoterMultiResponse types.RemoterMultiResponse = types.RemoterMultiResponse{}
var services map[string][]types.Service = map[string][]types.Service{}

const (
	URL = "http://%s:%d%s"
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
	services = remoterRes.Data
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
	fmt.Println(services)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			_ = <-ticker.C
			err = Sync(cfg.Names)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(services)
		}
	}()
	return nil
}

func DoRequest(name, method, path string, request interface{}, res interface{}) error {
	service, err := getService(name)
    if err != nil {
        return err
    }
	client := &http.Client{}
	url := fmt.Sprintf(URL, service.Host, service.Port, path)
	body, _ := json.Marshal(request)
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, res)
}

func getService(name string) (types.Service, error) {
	servs, ok := services[name]
	if !ok {
		return types.Service{}, errors.New("undefined service")
	}
	return servs[rand.Int31n(int32(len(servs)))], nil
}
