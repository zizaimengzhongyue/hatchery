package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const (
	Name    = "Server-Manager"
	Version = "0.0.1"
)

type Service struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

var services = map[string][]Service{}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func Debug(v interface{}) {
	bts, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bts))
}

func CreateID() string {
	uniqID := uuid.NewV4()
	return uniqID.String()
}

func Register(res http.ResponseWriter, req *http.Request) {
	bts, _ := ioutil.ReadAll(req.Body)
	service := &Service{}
	_ = json.Unmarshal(bts, service)
	service.ID = CreateID()
	services[service.Name] = append(services[service.Name], *service)
	bts, _ = json.Marshal(Response{Status: 0, Msg: "ok", Data: *service})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func Cancel(res http.ResponseWriter, req *http.Request) {
	bts, _ := ioutil.ReadAll(req.Body)
	service := &Service{}
	_ = json.Unmarshal(bts, service)
	for i, v := range services[service.Name] {
		if v.ID == service.ID {
			services[service.Name] = append(services[service.Name][0:i], services[service.Name][i+1:]...)
		}
	}
	bts, _ = json.Marshal(Response{Status: 0, Msg: "ok"})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func Get(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	bts, _ := json.Marshal(Response{Status: 0, Msg: "ok", Data: services[name]})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/cancel", Cancel)
	http.HandleFunc("/get", Get)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
