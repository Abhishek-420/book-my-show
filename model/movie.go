package model

type Movie struct {
	MovieId   string `json:"movieId"`
	MovieName string `json:"movieName"`
	Price     int    `json:"price"`
}

func NewMovie(movieId, movieName string, price int) *Movie {
	return &Movie{
		MovieId:   movieId,
		MovieName: movieName,
		Price:     price,
	}
}
