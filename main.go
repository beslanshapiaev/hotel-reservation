package main

import (
	"flag"

	"github.com/beslanshapiaev/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "application port")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1/")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
