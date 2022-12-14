package main

import (
	"example/realworld-api/src/controllers"
	config "example/realworld-api/src/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	// populateDb(cfg)

	router := gin.Default()

	router.GET("/api/users", controllers.GetUsers(cfg))
	router.POST("/api/users", controllers.RegisterUser(cfg))

	router.Run("0.0.0.0:" + cfg.PORT)
}

// func populateDb(cfg *config.Config) {
// 	cfg.DB.Users = append(cfg.DB.Users, schemas.User{Username: "user1"})
// 	cfg.DB.Users = append(cfg.DB.Users, schemas.User{Username: "user2"})
// 	cfg.DB.Users = append(cfg.DB.Users, schemas.User{Username: "user3"})
// }
