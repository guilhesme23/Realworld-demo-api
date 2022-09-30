package main

import (
	"fmt"

	config "example/realworld-api/src/internal"
)

func main() {
	config, _ := config.New()

	fmt.Println(config.PORT)
}
