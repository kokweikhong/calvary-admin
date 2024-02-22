package main

import (
	"github.com/joho/godotenv"
	"github.com/kokweikhong/calvary-admin/backend/internal/db"
	"github.com/kokweikhong/calvary-admin/backend/internal/routes"
)

func main() {
	godotenv.Load()

	postgresDB := db.NewPostgres()
	postgresDB.Connect()

	routes := routes.NewRoutes()
	router := routes.RegisterRoutes()
	routes.ListenAndServe(router, "8080")

	// mux := http.NewServeMux()
	// allow CORS

	// userRoutes := routes.NewUserRoutes()
	// userRoutes.RegisterRoutes(mux)
	// fileRoutes := routes.NewFileRoutes()
	// fileRoutes.RegisterRoutes(mux)

	// http.ListenAndServe(":8080", mux)
}
