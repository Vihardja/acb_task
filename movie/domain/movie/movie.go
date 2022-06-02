package movie

import (
	"acb_task/movie/model"
	"context"
)

type MovieRepository interface {
	AddMovie(ctx context.Context, req model.AddMovieRequest) (int64, error)
	GetMovies(ctx context.Context) ([]model.MovieData, error)
	GetMovieDetailByID(ctx context.Context, id int64) (*model.MovieData, error)
}
