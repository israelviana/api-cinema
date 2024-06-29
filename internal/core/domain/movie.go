package domain

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Duration string `json:"duration"`
}

type MovieTheaterShowing struct {
	MovieID        int `json:"movie_id"`
	MovieTheaterID int `json:"movie_theater_id"`
}
