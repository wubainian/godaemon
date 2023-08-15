package router

import (
	"github.com/wubainian/godaemon/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
