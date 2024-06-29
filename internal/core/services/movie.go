package services

import (
	"api-cinema/internal/core/domain"
	"api-cinema/internal/core/ports"
	"errors"
)

type movieService struct {
	movieRepository ports.MovieRepository
}

func NewMovieService(movieRepository ports.MovieRepository) *movieService {
	return &movieService{movieRepository: movieRepository}
}

func (srv *movieService) Get(id int) (domain.Movie, error) {
	movie, err := srv.movieRepository.GetMovie(id)
	if err != nil {
		return domain.Movie{}, errors.New("movie not found")
	}

	return movie, nil
}

func (srv *movieService) GetAll() ([]domain.Movie, error) {
	movies, err := srv.movieRepository.GetAllMovie()
	if err != nil {
		return []domain.Movie{}, errors.New("any movie not found")
	}

	return movies, nil
}

func (srv *movieService) Create(movie domain.Movie) (domain.Movie, error) {
	err := srv.movieRepository.SaveMovie(movie)
	if err != nil {
		return domain.Movie{}, errors.New("error to create movie")
	}

	return movie, nil
}

func (srv *movieService) Update(movie domain.Movie) (domain.Movie, error) {
	err := srv.movieRepository.UpdateMovie(movie)
	if err != nil {
		return domain.Movie{}, errors.New("error to update movie")
	}

	return movie, nil
}

func (srv *movieService) Delete(id int) error {
	err := srv.movieRepository.DeleteMovie(id)
	if err != nil {
		return errors.New("error to delete movie")
	}

	return nil
}
