package domain

import (
	"acb_task/movie/domain/movie"
	"acb_task/movie/model"
	"context"
	"database/sql"
)

type mysqlMovieRepo struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of movie.MovieRepository
func NewMysqlMovieRepository(db *sql.DB) movie.MovieRepository {
	return &mysqlMovieRepo{
		DB: db,
	}
}

func (m *mysqlMovieRepo) GetMovieDetailByID(ctx context.Context, id int64) (*model.MovieData, error) {

	query := `
	SELECT id, name, genre, release_year, production_house
	FROM movie
	WHERE id = ?
	`

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)
	result := new(model.MovieData)

	err = row.Scan(
		&result.ID,
		&result.Name,
		&result.Genre,
		&result.ReleaseYear,
		&result.ProductionHouse,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
