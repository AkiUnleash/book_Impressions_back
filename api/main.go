package main

import (
	"os"
	"srb/domain/database"
	"srb/domain/routing"

	"github.com/joho/godotenv"
)

// @title book_Impressions_back
// @version 1.0.0
// @description API of software to describe impressions of books.
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host Secret
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
