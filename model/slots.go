package model

type Slot struct {
	SlotID         string `json:"slotID"`
	SlotTime       string `json:"slotTime"`
	TheaterId      string `json:"theaterId"`
	MovieId        string `json:"movieId"`
	AvailableSeats int    `json:"AvailableSeats"`
}

func NewSlot(theaterID string, movieID string, slotID string, availableSeats int, slotTime string) *Slot {
	return &Slot{
		TheaterId:      theaterID,
		MovieId:        movieID,
		SlotID:         slotID,
		AvailableSeats: availableSeats,
		SlotTime:       slotTime,
	}
}
