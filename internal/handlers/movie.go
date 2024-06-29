package handlers

import (
	"api-cinema/internal/core/domain"
	"api-cinema/internal/core/ports"
	"api-cinema/pkg/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	movieService ports.MovieService
}

func NewMovieHandler(movieService ports.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

// FindMoviePerID godoc
// @Summary Find a movie by ID
// @Description Get details of a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 {object} domain.Movie
// @Failure 400 {object} domain.HTTPErrorResponse "Invalid ID"
// @Failure 404 {object} domain.HTTPErrorResponse "Movie not found"
// @Router /movies/{id} [get]
func (h *MovieHandler) FindMoviePerID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	movieID, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid id")
	}

	movier, err := h.movieService.Get(movieID)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "error to find movie per id: "+strconv.Itoa(movieID))
	}

	return ctx.Status(http.StatusOK).JSON(&movier)
}

// FindAllMovie godoc
// @Summary Get all movies
// @Description Retrieve a list of all movies
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {array} domain.Movie
// @Failure 404 {object} domain.HTTPErrorResponse
// @Router /movies [get]
func (h *MovieHandler) FindAllMovie(ctx *fiber.Ctx) error {
	listOfMovieTheater, err := h.movieService.GetAll()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "error to get all movies")
	}

	return ctx.Status(http.StatusOK).JSON(&listOfMovieTheater)
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie in the theater
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} domain.Movie
// @Failure 400 {object} domain.HTTPErrorResponse "Invalid ID"
// @Failure 404 {object} domain.HTTPErrorResponse "Movie not found"
// @Router /movies/{id} [post]
func (h *MovieHandler) CreateMovie(ctx *fiber.Ctx) error {
	var request domain.Movie

	err := json.Unmarshal(ctx.Body(), &request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to read received body")
	}

	movie, err := h.movieService.Create(request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to create movie")
	}

	return ctx.Status(http.StatusOK).JSON(&movie)
}

// UpdateMovie godoc
// @Summary Update an existing movie
// @Description Update an existing movie in the theater
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body domain.Movie true "Movie to update"
// @Success 200 {object} domain.Movie
// @Failure 400 {object} domain.HTTPErrorResponse
// @Router /movies [put]
func (h *MovieHandler) UpdateMovie(ctx *fiber.Ctx) error {
	var request domain.Movie

	err := json.Unmarshal(ctx.Body(), &request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to read received body")
	}

	movie, err := h.movieService.Update(request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to update movie")
	}

	return ctx.Status(http.StatusOK).JSON(&movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} string
// @Failure 400 {object} domain.HTTPErrorResponse
// @Router /movies [delete]
func (h *MovieHandler) DeleteMovie(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	movieID, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid id")
	}

	err = h.movieService.Delete(movieID)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "error to delete movie:  "+strconv.Itoa(movieID))
	}

	return ctx.Status(http.StatusOK).JSON("movie deleted with success")
}
