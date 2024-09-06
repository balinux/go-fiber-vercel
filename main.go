package main

import (
	"go-fiber-vercel/api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables dari file .env

	app := fiber.New()

	// Menambahkan routing
	routes.RootRoute(app)
	routes.V1Route(app)
	routes.V2Route(app)

	// Menentukan port dari environment variable atau default 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Menjalankan server
	log.Printf("Starting server on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
