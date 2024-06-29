CREATE DATABASE IF NOT EXISTS cinema_db;

CREATE TABLE IF NOT EXISTS movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    director VARCHAR(255) NOT NULL,
    duration VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS movie_theaters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    theater_number INT NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE movie_theater_showings (
    movie_id INT NOT NULL,
    theater_id INT NOT NULL,
    PRIMARY KEY (movie_id, theater_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (theater_id) REFERENCES movie_theaters(id)
);