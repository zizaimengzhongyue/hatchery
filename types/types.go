package types

type Service struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type RemoterMultiResponse struct {
	Status int                  `json:"stauts"`
	Msg    string               `json:"msg"`
	Data   map[string][]Service `json:"data"`
}

type Config struct {
	Host  string
	Port  int
	Names []string
}
