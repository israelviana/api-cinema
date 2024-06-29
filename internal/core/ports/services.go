package ports

import "api-cinema/internal/core/domain"

type MovieService interface {
	Get(id int) (domain.Movie, error)
	GetAll() ([]domain.Movie, error)
	Create(movie domain.Movie) (domain.Movie, error)
	Update(movie domain.Movie) (domain.Movie, error)
	Delete(id int) error
}

type MovieTheaterService interface {
	Get(id int) (domain.MovieTheater, error)
	GetAll() ([]domain.MovieTheater, error)
	Create(theater domain.MovieTheater) (domain.MovieTheater, error)
	Update(theater domain.MovieTheater) (domain.MovieTheater, error)
	AssociateMovieAtMovieTheater(movieTheaterShowing domain.MovieTheaterShowing) error
}
