package repository

import (
	"book-my-show/model"
)

type BookMyShowRepo struct {
	movieRepo      map[string]*model.Movie
	slotRepo       map[string]*model.Slot
	theaterRepo    map[string]*model.Theater
	movieTheaterID map[string][]*model.Slot
	movies         map[string][]*model.Theater
	booking        map[string]*model.Booking
}

type BookMyShowRepoInterface interface {
	GetAllMovies() []*model.Movie
	GetAllTheatersForMovie(movieID string) []*model.Theater
	GetAllSlotsByMovieIDTheaterID(movieID string, theaterID string) []*model.Slot
	AddMovie(movieID string, movieName string, price int) *model.Movie
	AddTheater(theaterID string, theaterName string, seats int) *model.Theater
	AddSlotForMovie(slotID string, slotTime string, movieId string, theaterID string, availableSeats int) *model.Slot
}

func NewBookMyShowRepo() BookMyShowRepoInterface {
	return &BookMyShowRepo{
		movieRepo:      make(map[string]*model.Movie),
		slotRepo:       make(map[string]*model.Slot),
		theaterRepo:    make(map[string]*model.Theater),
		movies:         make(map[string][]*model.Theater),
		movieTheaterID: make(map[string][]*model.Slot),
		booking:        make(map[string]*model.Booking),
	}
}

func (b *BookMyShowRepo) GetAllMovies() []*model.Movie {

	movies := make([]*model.Movie, 0)

	for _, movie := range b.movieRepo {
		movies = append(movies, movie)
	}
	return movies
}

func (b *BookMyShowRepo) GetAllTheatersForMovie(movieID string) []*model.Theater {
	return b.movies[movieID]
}

func (b *BookMyShowRepo) GetAllSlotsByMovieIDTheaterID(movieId string, theaterID string) []*model.Slot {
	key := movieId + theaterID
	return b.movieTheaterID[key]
}

func (b *BookMyShowRepo) AddMovie(movieId string, movieName string, price int) *model.Movie {

	movie := model.NewMovie(movieId, movieName, price)
	b.movieRepo[movieId] = movie
	return movie

}

func (b *BookMyShowRepo) AddTheater(theaterId string, theaterName string, seats int) *model.Theater {

	theater := model.NewTheater(theaterId, theaterName, seats)
	b.theaterRepo[theaterId] = theater
	return theater

}

func (b *BookMyShowRepo) AddSlotForMovie(slotId string, slotTime string, movieId string, theaterID string, availableSeats int) *model.Slot {

	slot := model.NewSlot(theaterID, movieId, slotId, availableSeats, slotTime)
	b.slotRepo[slotId] = slot

	moviedTheaterID := movieId + theaterID

	slots, ok := b.movieTheaterID[moviedTheaterID]
	if ok {
		slots = append(slots, slot)
		b.movieTheaterID[moviedTheaterID] = slots
	} else {

		newSlots := make([]*model.Slot, 0)
		newSlots = append(newSlots, slot)
		b.movieTheaterID[moviedTheaterID] = newSlots

	}

	Theaters, ok1 := b.movies[movieId]

	if ok1 {
		Theaters = append(Theaters, b.theaterRepo[theaterID])
		b.movies[movieId] = Theaters
	} else {

		newTheaterList := make([]*model.Theater, 0)
		newTheaterList = append(newTheaterList, b.theaterRepo[theaterID])
		b.movies[movieId] = newTheaterList

	}
	return slot

}

func (b *BookMyShowRepo) CreateBooking(movieId string, theaterId string, slotId string, reservedSeats int) *model.Booking {

}
