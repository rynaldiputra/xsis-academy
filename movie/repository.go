package movie

import (
	"log"
	"theater/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetMovie() ([]entities.Movie, error)
	GetMovieById(ID int) (entities.Movie, error)
	StoreMovie(movie entities.Movie) (entities.Movie, error)
	UpdateMovie(input MovieInput, ID int) (entities.Movie, error)
	DestroyMovie(ID int) (entities.Movie, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMovie() ([]entities.Movie, error) {
	var mb []entities.Movie

	if err := r.db.Where("deleted_at IS NULL").Find(&mb).Error; err != nil {
		return mb, err
	}

	return mb, nil
}

func (r *repository) GetMovieById(ID int) (entities.Movie, error) {
	var mb entities.Movie

	if err := r.db.Where("id = ?", ID).Where("deleted_at IS NULL").Find(&mb).Error; err != nil {
		return mb, err
	}

	return mb, nil
}

func (r *repository) StoreMovie(movie entities.Movie) (entities.Movie, error) {
	err := r.db.Create(&movie).Error

	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (r *repository) UpdateMovie(input MovieInput, ID int) (entities.Movie, error) {
	var mb entities.Movie

	if err := r.db.Where("id = ?", ID).Where("deleted_at IS NULL").Find(&mb).Error; err != nil {
		return mb, err
	}

	mb.Title = input.Title
	mb.Description = input.Description
	mb.Rating = input.Rating
	mb.Image = input.Image

	err := r.db.Updates(&mb).Error

	if err != nil {
		return entities.Movie{}, err
	}

	return mb, nil
}

func (r *repository) DestroyMovie(ID int) (entities.Movie, error) {
	var mb entities.Movie

	err := r.db.Delete(&mb, ID).Where("deleted_at IS NULL").Error

	if err != nil {
		return entities.Movie{}, err
	}

	log.Print(mb)

	return mb, nil
}
