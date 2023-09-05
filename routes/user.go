package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRoutes(server *gin.RouterGroup) {

	Group := server.Group("/public/v1/")
	{
		Group.POST("/register", cr.userRegister)
		Group.POST("/login", cr.userLogin)
	}
}
