package app

import (
	"ms/app/controllers/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(server *gin.Engine) {
	server.GET("/user/list", user.List)
	server.POST("/user/add", user.Add)
	server.GET("/user/:id", user.Detail)
}
