package model

import "time"

type Metadata struct {
	ID        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type MovieData struct {
	Metadata
	Name            string `db:"name" json:"name"`
	Genre           string `db:"genre" json:"genre"`
	ReleaseYear     int    `db:"release_year" json:"release_year"`
	ProductionHouse string `db:"production_house" json:"production_house"`
}

type AddMovieRequest struct {
	Name            string `json:"name"`
	Genre           string `json:"genre"`
	ReleaseYear     int    `json:"release_year"`
	ProductionHouse string `json:"production_house"`
}
