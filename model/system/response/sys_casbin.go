package response

import (
	"github.com/wubainian/godaemon/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
