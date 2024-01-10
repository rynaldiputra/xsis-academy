package movie

import (
	"theater/entities"
)

type Service interface {
	GetMovies() ([]entities.Movie, error)
	GetMovie(ID int) (entities.Movie, error)
	CreateMovie(input MovieInput) (entities.Movie, error)
	EditMovie(input MovieInput, ID int) (entities.Movie, error)
	DeleteMovie(ID int) (entities.Movie, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMovies() ([]entities.Movie, error) {
	movie, err := s.repository.GetMovie()

	if err != nil {
		return []entities.Movie{}, err
	}

	return movie, err
}

func (s *service) GetMovie(ID int) (entities.Movie, error) {
	movie, err := s.repository.GetMovieById(ID)

	if err != nil {
		return entities.Movie{}, err
	}

	return movie, err
}

func (s *service) CreateMovie(input MovieInput) (entities.Movie, error) {
	var movie entities.Movie

	movie.Title = input.Title
	movie.Description = input.Description
	movie.Rating = input.Rating
	// movie.Image = input.Image

	newBook, err := s.repository.StoreMovie(movie)

	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) EditMovie(input MovieInput, ID int) (entities.Movie, error) {
	movie, err := s.repository.UpdateMovie(input, ID)

	if err != nil {
		return entities.Movie{}, err
	}

	return movie, err
}

func (s *service) DeleteMovie(ID int) (entities.Movie, error) {
	movie, err := s.repository.DestroyMovie(ID)

	if err != nil {
		return entities.Movie{}, err
	}

	return movie, err
}
