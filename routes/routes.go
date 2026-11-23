package routes

import (
	"book-my-show/handler"
	"net/http"
)

func RegisterRoutes(h *handler.BookMyShowHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /movies", h.AddMovie)
	mux.HandleFunc("GET /movies", h.GetAllMovies)
	mux.HandleFunc("GET /movies/{movieId}", h.GetAllSlotsForMovie) // optional query: ?theaterId=T1

	mux.HandleFunc("POST /theaters", h.AddTheater)
	mux.HandleFunc("POST /slots", h.AddSlotForMovie)
	mux.HandleFunc("POST /bookings", h.BookMovie)

	return mux
}
