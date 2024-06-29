package ports

import "api-cinema/internal/core/domain"

type MovieRepository interface {
	GetMovie(id int) (domain.Movie, error)
	GetAllMovie() ([]domain.Movie, error)
	SaveMovie(movie domain.Movie) error
	UpdateMovie(movie domain.Movie) error
	DeleteMovie(id int) error
}

type MovieTheaterRepository interface {
	GetMovieTheater(id int) (domain.MovieTheater, error)
	GetAllMovieTheater() ([]domain.MovieTheater, error)
	SaveMovieTheater(movie domain.MovieTheater) error
	UpdateMovieTheater(movie domain.MovieTheater) error
	AssociateMovieAtMovieTheater(movieTheaterShowing domain.MovieTheaterShowing) error
}
