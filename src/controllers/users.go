package controllers

import (
	config "example/realworld-api/src/internal"
	"example/realworld-api/src/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(cf *config.Config) gin.HandlerFunc {
	users := cf.DB.Users
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
			db := cf.DB
			db.Users = append(db.Users, schemas.User{
				Username: body.User.Username,
				Password: body.User.Password,
				Email:    body.User.Email,
			})

			ctx.IndentedJSON(http.StatusCreated, gin.H{
				"user": db.Users[len(db.Users)-1],
			})
		} else {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
		}
	}
}
