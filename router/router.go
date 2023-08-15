package router

import "github.com/gin-gonic/gin"

var RouterPrivateInit func(group *gin.RouterGroup)
var RouterPublicInit func(group *gin.RouterGroup)

func RegisterAllPrivateRouters(group *gin.RouterGroup) {
	// RouterInit = func(group *gin.RouterGroup) {
	// }
	RouterPrivateInit(group)
}

func RegisterAllPublicRouters(group *gin.RouterGroup) {
	// RouterInit = func(group *gin.RouterGroup) {
	// }
	RouterPublicInit(group)
}
