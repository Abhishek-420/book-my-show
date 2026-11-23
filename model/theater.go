package model

type Theater struct {
	TheaterID   string `json:"theaterId"`
	TheaterName string `json:"theaterName"`
	Seats       int    `json:"seats"`
}

func NewTheater(theaterId, theaterName string, seats int) *Theater {
	return &Theater{
		TheaterID:   theaterId,
		TheaterName: theaterName,
		Seats:       seats,
	}
}
