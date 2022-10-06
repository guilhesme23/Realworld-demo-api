package controllers

import (
	config "example/realworld-api/src/internal"
	"example/realworld-api/src/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []schemas.User{}

func GetUsers(cf *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, users)
	}
}

func RegisterUser(cf *config.Config) gin.HandlerFunc {
	db := cf.DB
	return func(ctx *gin.Context) {
		var body struct {
			User schemas.User `json:"user" binding:"required"`
		}

		if ctx.Bind(&body) == nil {
			var user schemas.User = body.User
			res := db.Create(&user)
			if res.Error == nil {
				ctx.IndentedJSON(http.StatusCreated, gin.H{
					"user": user,
				})
			} else {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"message": "error",
				})
			}
		} else {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
		}
	}
}
