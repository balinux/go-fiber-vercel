package handler

import (
	"go-fiber-vercel/routes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// Handler is the main entry point of the application.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// Building the fiber application
func handler() http.HandlerFunc {
	app := fiber.New()

	// menambahkan routing
	routes.RootRoute(app)
	routes.V1Route(app)
	routes.V2Route(app)

	return adaptor.FiberApp(app)
}
