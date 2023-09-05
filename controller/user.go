package main

import (
	models "design-system/model"
	"design-system/utils"
	"net/http"

	"go.elastic.co/apm/v2"

	"github.com/gin-gonic/gin"
)

// var (
//     userService = services.NewUserService()
// )

func userRegister(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "register", "controller")
	defer span.End()
	ctx = utils.SetContext(c, ctx)

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error Invalid ReqBody",
		})
		return
	}

	userID, err := service.RegisterUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered", "userID": userID})
}

func userLogin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "login", "controller")
	defer span.End()
	ctx = utils.SetContext(c, ctx)

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error Invalid ReqBody",
		})
		return
	}

	foundUser, err := userService.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated", "user": foundUser})
}
