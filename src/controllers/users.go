package controllers

import (
	config "example/realworld-api/src/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(cf *config.Config) gin.HandlerFunc {
	users := cf.DB.Users
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, users)
	}
}
