package usecase

import (
	"acb_task/movie/domain/movie"
	"acb_task/movie/model"
	"context"
	"fmt"
)

type MovieUsecase interface {
	AddMovie(ctx context.Context, req model.AddMovieRequest) (string, error)
	GetMovies(ctx context.Context) ([]model.MovieData, error)
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

func (m *movieUsecase) AddMovie(ctx context.Context, req model.AddMovieRequest) (string, error) {
	id, err := m.movieRepo.AddMovie(ctx, req)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully add new movie with ID: %d", id), nil
}

func (m *movieUsecase) GetMovies(ctx context.Context) ([]model.MovieData, error) {
	result, err := m.movieRepo.GetMovies(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *movieUsecase) GetMovieDetail(ctx context.Context, id int64) (*model.MovieData, error) {
	result, err := m.movieRepo.GetMovieDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
