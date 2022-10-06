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
	return func(ctx *gin.Context) {
		var body struct {
			User schemas.User `json:"user" binding:"required"`
		}

		if ctx.Bind(&body) == nil {
			users = append(users, schemas.User{
				Username: body.User.Username,
				Password: body.User.Password,
				Email:    body.User.Email,
			})

			ctx.IndentedJSON(http.StatusCreated, gin.H{
				"user": users[len(users)-1],
			})
		} else {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
		}
	}
}
