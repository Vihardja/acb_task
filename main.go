package main

import (
	domain "acb_task/movie/domain/database"
	http "acb_task/movie/http"
	usecase "acb_task/movie/usecase"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbUser                 = "root"
		dbPwd                  = ""
		instanceConnectionName = "studious-spot-352021:us-central1:acbmovie1"
		dbName                 = "v1_movie"
	)

	dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", dbUser, dbPwd, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalln("sql.Open: %v", err)
	}
	err = dbPool.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbPool.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	mvRepo := domain.NewMysqlMovieRepository(dbPool)
	mu := usecase.NewMovieUsecase(mvRepo)
	http.NewRequestHandler(mu)
}
