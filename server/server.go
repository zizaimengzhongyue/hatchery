package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/zizaimengzhongyue/hatchery/logging"
	"github.com/zizaimengzhongyue/hatchery/types"
)

var services = map[string][]types.Service{}

func Debug(v interface{}) {
	bts, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bts))
}

func CreateID() string {
	return uuid.NewV4().String()
}

func isRegistered(serv *types.Service) bool {
	for _, v := range services[serv.Name] {
		if v.Host == serv.Host && v.Port == serv.Port {
			return true
		}
	}
	return false
}

func Register(res http.ResponseWriter, req *http.Request) {
	bts, _ := ioutil.ReadAll(req.Body)
	logging.Get("hatchery").Println(string(bts))
	service := &types.Service{}
	_ = json.Unmarshal(bts, service)
	if isRegistered(service) {
		bts, _ = json.Marshal(types.Response{Status: 1, Msg: "ok", Data: "have registered"})
		goto res
	}
	service.ID = CreateID()
	services[service.Name] = append(services[service.Name], *service)
	bts, _ = json.Marshal(types.Response{Status: 0, Msg: "ok", Data: *service})
res:
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func Cancel(res http.ResponseWriter, req *http.Request) {
	bts, _ := ioutil.ReadAll(req.Body)
	service := &types.Service{}
	_ = json.Unmarshal(bts, service)
	for i, v := range services[service.Name] {
		if v.ID == service.ID {
			services[service.Name] = append(services[service.Name][0:i], services[service.Name][i+1:]...)
		}
	}
	bts, _ = json.Marshal(types.Response{Status: 0, Msg: "ok"})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func Get(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	bts, _ := json.Marshal(types.Response{Status: 0, Msg: "ok", Data: services[name]})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func MultiGet(res http.ResponseWriter, req *http.Request) {
	names := strings.Split(req.URL.Query().Get("names"), ",")
	ans := map[string][]types.Service{}
	for _, name := range names {
		ans[name] = services[name]
	}
	bts, _ := json.Marshal(types.Response{Status: 0, Msg: "ok", Data: ans})
	res.Header().Add("Content-Type", "application/json")
	res.Write(bts)
}

func Init() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/cancel", Cancel)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/multiGet", MultiGet)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
