package main

import (
	"api-cinema/internal/server"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// @title API-CINEMA
// @description API for cinema operations.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()
	server.InitServer(app)
	fmt.Println("ta rodando")
}
