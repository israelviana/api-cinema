package repositories

import (
	"api-cinema/internal/core/domain"
	"api-cinema/pkg/db/mysql"
	"database/sql"
)

type movieTheaterRepository struct {
	mysqldb *sql.DB
}

func NewMovieTheaterRepository() *movieTheaterRepository {
	return &movieTheaterRepository{mysqldb: mysql.NewMySQLConnection()}
}

func (db movieTheaterRepository) GetMovieTheater(id int) (domain.MovieTheater, error) {
	query := `SELECT mt.id as movie_theater_id, mt.theater_number, mt.description, COALESCE(mtv.movie_id, 0), COALESCE(mv.director, ''), COALESCE(mv.duration, ''), COALESCE(mv.title, '') FROM movie_theaters as mt
    LEFT JOIN movie_theater_showings as mtv ON mt.id = mtv.theater_id
    LEFT JOIN movies as mv ON mtv.movie_id = mv.id
                                                            WHERE mt.id=?`

	var movieTheater domain.MovieTheater
	rows, err := db.mysqldb.Query(query, id)
	if err != nil {
		return movieTheater, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var movie domain.Movie
			if err = rows.Scan(&movieTheater.ID, &movieTheater.TheaterNumber, &movieTheater.Description, &movie.ID, &movie.Director, &movie.Duration, &movie.Title); err != nil {
				return movieTheater, err
			}

			movieTheater.Movies = append(movieTheater.Movies, movie)
		}
	}

	return movieTheater, nil
}

func (db movieTheaterRepository) GetAllMovieTheater() ([]domain.MovieTheater, error) {
	query := `SELECT mt.id as movie_theater_id, mt.theater_number, mt.description, COALESCE(mtv.movie_id, 0), COALESCE(mv.director, ''), COALESCE(mv.duration, ''), COALESCE(mv.title, '') FROM movie_theaters as mt
    LEFT JOIN movie_theater_showings as mtv ON mt.id = mtv.theater_id
    LEFT JOIN movies as mv ON mtv.movie_id = mv.id;`

	var allMovieTheater []domain.MovieTheater
	rows, err := db.mysqldb.Query(query)
	if err != nil {
		return allMovieTheater, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var movieTheater domain.MovieTheater
			var movie domain.Movie
			if err = rows.Scan(&movieTheater.ID, &movieTheater.TheaterNumber, &movieTheater.Description, &movie.ID, &movie.Director, &movie.Duration, &movie.Title); err != nil {
				return allMovieTheater, err
			}

			if movie.ID != 0 {
				movieTheater.Movies = append(movieTheater.Movies, movie)
			}
			allMovieTheater = append(allMovieTheater, movieTheater)
		}
	}

	return allMovieTheater, nil
}

func (db movieTheaterRepository) SaveMovieTheater(movieTheater domain.MovieTheater) error {
	query := "INSERT INTO movie_theaters (theater_number, description) VALUES (?, ?)"

	_, err := db.mysqldb.Exec(query, movieTheater.TheaterNumber, movieTheater.Description)
	if err != nil {
		return err
	}

	return nil
}

func (db movieTheaterRepository) UpdateMovieTheater(movieTheater domain.MovieTheater) error {
	query := "UPDATE movie_theaters SET theater_number = ?, description = ? WHERE id = ?"

	_, err := db.mysqldb.Exec(query, movieTheater.TheaterNumber, movieTheater.Description, movieTheater.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db movieTheaterRepository) AssociateMovieAtMovieTheater(movieTheaterShowing domain.MovieTheaterShowing) error {
	query := "INSERT INTO movie_theater_showings (movie_id, theater_id) VALUES (?, ?)"

	_, err := db.mysqldb.Exec(query, movieTheaterShowing.MovieID, movieTheaterShowing.MovieTheaterID)
	if err != nil {
		return err
	}

	return nil
}
