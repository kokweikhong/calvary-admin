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
	userRoutes := routes.NewUserRoutes()
	userRoutes.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
