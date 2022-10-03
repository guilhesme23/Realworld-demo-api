package main

import (
	"net/http"

	"example/realworld-api/src/controllers"
	config "example/realworld-api/src/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.New()

	cfg.DB.Users = append(cfg.DB.Users, config.User{Username: "user1"})
	cfg.DB.Users = append(cfg.DB.Users, config.User{Username: "user2"})
	cfg.DB.Users = append(cfg.DB.Users, config.User{Username: "user3"})
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/users", controllers.GetUsers(cfg))

	router.Run("0.0.0.0:" + cfg.PORT)
}
