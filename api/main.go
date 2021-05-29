package main

import (
	"os"
	"srb/domain/database"
	"srb/domain/routing"

	"github.com/joho/godotenv"
)

// @title JWT-login-example
// @version 1.0.0
// @description API for login processing using JWT. Developed in Go langage.

// @license.name MIT
// @license.url https://github.com/tcnksm/tool/blob/master/LICENCE

// @host http://localhost:8081
// @BasePath /
func main() {
	if os.Getenv("ENVIROMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}

	database.Connet()

	routing.Routing()

}
