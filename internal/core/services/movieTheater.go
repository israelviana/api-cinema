package services

import (
	"api-cinema/internal/core/domain"
	"api-cinema/internal/core/ports"
	"errors"
)

type movieTheaterService struct {
	movieTheaterRepository ports.MovieTheaterRepository
}

func NewMovieTheaterService(movieTheaterRepository ports.MovieTheaterRepository) *movieTheaterService {
	return &movieTheaterService{movieTheaterRepository: movieTheaterRepository}
}

func (srv *movieTheaterService) Get(id int) (domain.MovieTheater, error) {
	movieTheater, err := srv.movieTheaterRepository.GetMovieTheater(id)
	if err != nil {
		return domain.MovieTheater{}, errors.New("movieTheater not found")
	}

	return movieTheater, nil
}

func (srv *movieTheaterService) GetAll() ([]domain.MovieTheater, error) {
	movies, err := srv.movieTheaterRepository.GetAllMovieTheater()
	if err != nil {
		return []domain.MovieTheater{}, errors.New("any movieTheater not found")
	}

	return movies, nil
}

func (srv *movieTheaterService) Create(movieTheater domain.MovieTheater) (domain.MovieTheater, error) {
	err := srv.movieTheaterRepository.SaveMovieTheater(movieTheater)
	if err != nil {
		return domain.MovieTheater{}, errors.New("error to create movieTheater")
	}

	return movieTheater, nil
}

func (srv *movieTheaterService) Update(movieTheater domain.MovieTheater) (domain.MovieTheater, error) {
	err := srv.movieTheaterRepository.UpdateMovieTheater(movieTheater)
	if err != nil {
		return domain.MovieTheater{}, errors.New("error to update movieTheater")
	}

	return movieTheater, nil
}

func (srv *movieTheaterService) AssociateMovieAtMovieTheater(movieTheaterShowing domain.MovieTheaterShowing) error {
	err := srv.movieTheaterRepository.AssociateMovieAtMovieTheater(movieTheaterShowing)
	if err != nil {
		return errors.New("error to associate movie at movie theater")
	}

	return nil
}
