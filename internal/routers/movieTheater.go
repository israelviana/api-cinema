package routers

import (
	"api-cinema/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MapMovieTheaterRouters(fiber fiber.Router, handler *handlers.MovieTheaterHandler) {
	fiber.Get("movie_theaters/:id", handler.FindMovieTheaterPerID)
	fiber.Get("movie_theaters", handler.FindAllMovieTheaters)
	fiber.Post("movie_theaters", handler.CreateMovieTheater)
	fiber.Put("movie_theaters", handler.UpdateMovieTheater)
	fiber.Post("associate", handler.AssociateMovieAtMovieTheater)
}
