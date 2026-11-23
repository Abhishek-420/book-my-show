package handler

import (
	"book-my-show/service"
	"encoding/json"
	"net/http"
)

type BookMyShowHandler struct {
	service service.BookMyShowService
}

func NewBookMyShowHandler(service *service.BookMyShowService) *BookMyShowHandler {
	return &BookMyShowHandler{
		service: *service,
	}
}

type addMovieRequest struct {
	MovieID   string `json:"movieId"`
	MovieName string `json:"movieName"`
	Price     int    `json:"price"`
}

type addTheaterRequest struct {
	TheaterID   string `json:"theaterId"`
	TheaterName string `json:"theaterName"`
	Seats       int    `json:"seats"`
}

type addSlotRequest struct {
	SlotId         string `json:"slotID"`
	SlotTime       string `json:"slotTime"`
	MovieID        string `json:"movieId"`
	TheaterID      string `json:"theaterId"`
	AvailableSeats int    `json:"availableSeats"`
}

type bookMovieRequest struct {
	MovieID       string `json:"movieId"`
	TheaterID     string `json:"theaterId"`
	SlotID        string `json:"slotId"`
	ReservedSeats int    `json:"reservedSeats"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *BookMyShowHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	var req addMovieRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	movie := h.service.AddMovie(req.MovieID, req.MovieName, req.Price)

	writeJSON(w, http.StatusCreated, movie)

}

func (h *BookMyShowHandler) AddTheater(w http.ResponseWriter, r *http.Request) {
	var req addTheaterRequest
	json.NewDecoder(r.Body).Decode(&req)

	theater := h.service.AddTheater(req.TheaterID, req.TheaterName, req.Seats)
	writeJSON(w, http.StatusCreated, theater)
}

func (h *BookMyShowHandler) AddSlotForMovie(w http.ResponseWriter, r *http.Request) {
	var req addSlotRequest
	json.NewDecoder(r.Body).Decode(&req)
	slot := h.service.AddSlotForMovie(req.SlotId, req.SlotTime, req.MovieID, req.TheaterID, req.AvailableSeats)
	writeJSON(w, http.StatusCreated, slot)
}

func (h *BookMyShowHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := h.service.GetAllMovies()
	writeJSON(w, http.StatusOK, movies)
}

func (h *BookMyShowHandler) GetAllSlotsForMovie(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("movieId")

	allDetails := h.service.GetAvaiableSlotsForMovie(movieId)
	writeJSON(w, http.StatusOK, allDetails)

}

func (h *BookMyShowHandler) BookMovie(w http.ResponseWriter, r *http.Request) {

	var req bookMovieRequest
	json.NewDecoder(r.Body).Decode(&req)

	booking := h.service.CreateBooking(req.MovieID, req.TheaterID, req.SlotID, req.ReservedSeats)
	writeJSON(w, http.StatusCreated, booking)

}
