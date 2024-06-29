package server

import (
	_ "api-cinema/docs"
	"api-cinema/internal/core/services"
	"api-cinema/internal/handlers"
	"api-cinema/internal/repositories"
	"api-cinema/internal/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func InitServer(app *fiber.App) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	//Init Services
	movieRepository := repositories.NewMovieRepository()
	movieTheaterRepository := repositories.NewMovieTheaterRepository()

	movieService := services.NewMovieService(movieRepository)
	movieTheaterService := services.NewMovieTheaterService(movieTheaterRepository)

	//InitHandlers
	movieHandler := handlers.NewMovieHandler(movieService)
	movieTheaterHandler := handlers.NewMovieTheaterHandler(movieTheaterService)

	//cors
	crs := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, PUT, OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
	})
	app.Use(crs)

	health := app.Group("/health")
	cinemaGroup := app.Group("/")

	routers.MapMovieRouters(cinemaGroup, movieHandler)
	routers.MapMovieTheaterRouters(cinemaGroup, movieTheaterHandler)

	app.Get("/swagger/*", swagger.HandlerDefault)
	health.Get("", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{"status": "OK"})
	})

	if err := app.Listen(os.Getenv("port_application")); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}
