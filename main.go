package main

import (
	"fmt"
	"net/http"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/gql"
	"scriptura/scriptura-api/router"

	"github.com/joho/godotenv"
)

// Scriptura
//
//	@title			Scriptura API
//	@version		1.0
//	@description	API for accessing Bible scriptures
//	@termsOfService	http://scriptura.dev/terms/
//	@contact.name	API Support
//	@contact.url	http://scriptura.dev/support
//	@contact.email	support@scriptura.dev
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/api/v1
func main() {
	godotenv.Load()
	db.Connect()

	r := router.NewRouter()

	err := gql.CreateSchema()
	if err != nil {
		panic(err)
	}

	fmt.Println("Scriptura 📜 is up on port 3000")
	panic(http.ListenAndServe(":3000", r))
}
