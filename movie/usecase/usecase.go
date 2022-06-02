package usecase

import (
	"acb_task/movie/domain/movie"
	"acb_task/movie/model"
	"context"
)

type MovieUsecase interface {
	//AddMovie(searchWord string, pagination int) ([]byte, error)
	GetMovieDetail(ctx context.Context, id int64) (*model.MovieData, error)
}

type movieUsecase struct {
	movieRepo movie.MovieRepository
}

func NewMovieUsecase(mv movie.MovieRepository) MovieUsecase {
	return &movieUsecase{
		movieRepo: mv,
	}
}

func (m *movieUsecase) GetMovieDetail(ctx context.Context, id int64) (*model.MovieData, error) {
	result, err := m.movieRepo.GetMovieDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
