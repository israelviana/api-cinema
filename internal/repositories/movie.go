package repositories

import (
	"api-cinema/internal/core/domain"
	"api-cinema/pkg/db/mysql"
	"database/sql"
)

type movieRepository struct {
	mysqldb *sql.DB
}

func NewMovieRepository() *movieRepository {
	return &movieRepository{mysqldb: mysql.NewMySQLConnection()}
}

func (db movieRepository) GetMovie(id int) (domain.Movie, error) {
	query := `SELECT * FROM movies WHERE id=?`

	var movie domain.Movie
	err := db.mysqldb.QueryRow(query, id).Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Duration)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (db movieRepository) GetAllMovie() ([]domain.Movie, error) {
	query := `SELECT * FROM movies`

	var movies []domain.Movie
	rows, err := db.mysqldb.Query(query)
	if err != nil {
		return movies, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var movie domain.Movie
			if err = rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Duration); err != nil {
				return movies, err
			}

			movies = append(movies, movie)
		}
	}

	return movies, nil
}

func (db movieRepository) SaveMovie(movie domain.Movie) error {
	query := "INSERT INTO movies (title, director, duration) VALUES (?, ?, ?)"

	_, err := db.mysqldb.Exec(query, movie.Title, movie.Director, movie.Duration)
	if err != nil {
		return err
	}

	return nil
}

func (db movieRepository) UpdateMovie(movie domain.Movie) error {
	query := "UPDATE movies SET title = ?, director = ?, duration = ? WHERE id = ?"

	_, err := db.mysqldb.Exec(query, movie.Title, movie.Director, movie.Duration, movie.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db movieRepository) DeleteMovie(id int) error {
	query := "DELETE FROM movies WHERE id = ?"

	_, err := db.mysqldb.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
