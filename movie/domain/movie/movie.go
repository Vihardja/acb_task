package movie

import (
	"acb_task/movie/model"
	"context"
)

type MovieRepository interface {
	GetMovieDetailByID(ctx context.Context, id int64) (*model.MovieData, error)
}
