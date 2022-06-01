package usecase

import (
	"acb_task/movie/model"
)

type MovieUsecase interface {
	//AddMovie(searchWord string, pagination int) ([]byte, error)
	GetMovieDetail(id string) (*model.MovieData, error)
}

type movieUsecase struct{}

func NewMovieUsecase() MovieUsecase {
	return &movieUsecase{}
}

func (m *movieUsecase) GetMovieDetail(id string) (*model.MovieData, error) {

	return nil, nil
}
