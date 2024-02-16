package main

import (
	"fmt"
	"net/http"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/handler"
	"scriptura/scriptura-api/repository"
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
	db, _ := db.CreateDBConnection(10)

	repo := repository.NewAppRepository(db)
	handler := handler.NewAppHandlers(repo)
	r := router.NewAppRouter(&repo, &handler)

	fmt.Println("Scriptura 📜 is up on port 3000")
	panic(http.ListenAndServe(":3000", r))
}
