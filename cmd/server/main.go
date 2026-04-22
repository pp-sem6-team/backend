package main

import (
	"log"

	"github.com/pp-sem6-team/backend/internal/app"
)

/**
@title           Backend API
@version         1.0

@securityDefinitions.apikey BearerAuth
@in header
@name Authorization
@description Enter token as: Bearer <your_token>
*/

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
