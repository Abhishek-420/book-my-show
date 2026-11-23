package model

type Booking struct {
	BookingId     string `json:"bookingId"`
	MovieId       string `json:"movieId"`
	TheaterId     string `json:"theaterId"`
	SlotId        string `json:"slotId"`
	ReservedSeats int    `json:"reservedSeats"`
}

func NewBooking(bookingId, movieId, theaterId, slotId string, reservedSeats int) *model._Booking {

	return &Booking{
		BookingId:     bookingId,
		MovieId:       movieId,
		TheaterId:     theaterId,
		SlotId:        slotId,
		ReservedSeats: reservedSeats,
	}
}
