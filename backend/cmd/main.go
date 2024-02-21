package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kokweikhong/calvary-admin/backend/internal/db"
	"github.com/kokweikhong/calvary-admin/backend/internal/routes"
)

func main() {
	godotenv.Load()

	postgresDB := db.NewPostgres()
	postgresDB.Connect()

	mux := http.NewServeMux()
	// allow CORS

	userRoutes := routes.NewUserRoutes()
	userRoutes.RegisterRoutes(mux)
	fileRoutes := routes.NewFileRoutes()
	fileRoutes.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
