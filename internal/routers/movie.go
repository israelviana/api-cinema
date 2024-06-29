package routers

import (
	"api-cinema/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MapMovieRouters(fiber fiber.Router, handler *handlers.MovieHandler) {
	fiber.Get("movies/:id", handler.FindMoviePerID)
	fiber.Get("movies", handler.FindAllMovie)
	fiber.Post("movies", handler.CreateMovie)
	fiber.Put("movies", handler.UpdateMovie)
	fiber.Delete("movies/:id", handler.DeleteMovie)
}
