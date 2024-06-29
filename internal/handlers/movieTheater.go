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

type MovieTheaterHandler struct {
	movieTheaterService ports.MovieTheaterService
}

func NewMovieTheaterHandler(movieTheaterService ports.MovieTheaterService) *MovieTheaterHandler {
	return &MovieTheaterHandler{movieTheaterService: movieTheaterService}
}

// FindMovieTheaterPerID godoc
// @Summary Find a movie theater by ID
// @Description Get a movie theater by its ID
// @Tags movie_theaters
// @Accept json
// @Produce json
// @Param id path int true "Movie Theater ID"
// @Success 200 {object} domain.MovieTheater
// @Failure 400 {object} domain.HTTPErrorResponse
// @Failure 404 {object} domain.HTTPErrorResponse
// @Router /movie_theaters/{id} [get]
func (h *MovieTheaterHandler) FindMovieTheaterPerID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	movieTheaterID, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "invalid id")
	}

	movieTheater, err := h.movieTheaterService.Get(movieTheaterID)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "error to find movie theater per id: "+strconv.Itoa(movieTheaterID))
	}

	return ctx.Status(http.StatusOK).JSON(&movieTheater)
}

// FindAllMovieTheaters godoc
// @Summary Get all movie theaters
// @Description Retrieve a list of all movie theaters
// @Tags movie_theaters
// @Accept json
// @Produce json
// @Success 200 {array} domain.MovieTheater
// @Failure 404 {object} domain.HTTPErrorResponse
// @Router /movie_theaters [get]
func (h *MovieTheaterHandler) FindAllMovieTheaters(ctx *fiber.Ctx) error {
	listOfMovieTheater, err := h.movieTheaterService.GetAll()
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "error to get all movie theaters")
	}

	return ctx.Status(http.StatusOK).JSON(&listOfMovieTheater)
}

// CreateMovieTheater godoc
// @Summary Create a new movie theater
// @Description Create a new movie theater
// @Tags movie_theaters
// @Accept json
// @Produce json
// @Param movieTheater body domain.MovieTheater true "Movie Theater to create"
// @Success 200 {object} domain.MovieTheater
// @Failure 400 {object} domain.HTTPErrorResponse
// @Router /movie_theaters [post]
func (h *MovieTheaterHandler) CreateMovieTheater(ctx *fiber.Ctx) error {
	var request domain.MovieTheater

	err := json.Unmarshal(ctx.Body(), &request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to read received body")
	}

	movieTheater, err := h.movieTheaterService.Create(request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to create movie theater")
	}

	return ctx.Status(http.StatusOK).JSON(&movieTheater)
}

// UpdateMovieTheater godoc
// @Summary Update an existing movie theater
// @Description Update an existing movie theater
// @Tags movie_theaters
// @Accept json
// @Produce json
// @Param movieTheater body domain.MovieTheater true "Movie Theater to update"
// @Success 200 {object} domain.MovieTheater
// @Failure 400 {object} domain.HTTPErrorResponse
// @Router /movie_theaters [put]
func (h *MovieTheaterHandler) UpdateMovieTheater(ctx *fiber.Ctx) error {
	var request domain.MovieTheater

	err := json.Unmarshal(ctx.Body(), &request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to read received body")
	}

	movieTheater, err := h.movieTheaterService.Update(request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to update movie theater")
	}

	return ctx.Status(http.StatusOK).JSON(&movieTheater)
}

// AssociateMovieAtMovieTheater godoc
// @Summary Associate a movie at movie theater
// @Description Associate a movie at movie theater
// @Tags movie_theaters
// @Accept json
// @Produce json
// @Param MovieTheaterShowing body domain.MovieTheaterShowing true "Movie associate at movie theater"
// @Success 200 {object} string
// @Failure 400 {object} domain.HTTPErrorResponse
// @Router /associate [post]
func (h *MovieTheaterHandler) AssociateMovieAtMovieTheater(ctx *fiber.Ctx) error {
	var request domain.MovieTheaterShowing

	err := json.Unmarshal(ctx.Body(), &request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to read received body")
	}

	err = h.movieTheaterService.AssociateMovieAtMovieTheater(request)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "error to update movie theater")
	}

	return ctx.Status(http.StatusOK).JSON("movie is associated at the movie theater")
}
