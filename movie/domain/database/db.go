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

func (m *mysqlMovieRepo) AddMovie(ctx context.Context, req model.AddMovieRequest) (int64, error) {
	query := `
	INSERT INTO movie SET name=? , genre=? , release_year=?, production_house=? 
	`

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, req.Name, req.Genre, req.ReleaseYear, req.ProductionHouse)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (m *mysqlMovieRepo) GetMovies(ctx context.Context) ([]model.MovieData, error) {

	query := `
	SELECT id, name, genre, release_year, production_house
	FROM movie
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	result := make([]model.MovieData, 0)
	for rows.Next() {
		m := model.MovieData{}
		err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.Genre,
			&m.ReleaseYear,
			&m.ProductionHouse,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
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
