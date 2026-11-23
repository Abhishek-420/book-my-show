package service

import "book-my-show/repository"

type BookMyShowService struct {
	Service repository.BookMyShowRepoInterface
}

func NewBookMyShowService(repo repository.BookMyShowRepoInterface) *BookMyShowService {
	return &BookMyShowService{
		Service: repo,
	}
}

func (s *BookMyShowService) addMovie() {

	s.Service.AddMovie()
}
