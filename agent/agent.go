package agent

import (
    "github.com/zizaimengzhongyue/server-manager/service"
)

type Agent interface {
    Register(Service) (string, error)
    Cancel(Service) (string, error)
}
