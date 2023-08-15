package system

import "github.com/wubainian/godaemon/service"

type ApiGroup struct {
	JwtApi
	CasbinApi
}

var (
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
