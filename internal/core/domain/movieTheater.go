package domain

type MovieTheater struct {
	ID            int     `json:"id"`
	TheaterNumber int     `json:"theater_number"`
	Description   string  `json:"description"`
	Movies        []Movie `json:"movies"`
}
