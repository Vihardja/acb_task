package main

import (
	http "acb_task/movie/http"
	usecase "acb_task/movie/usecase"
)

func main() {
	mu := usecase.NewMovieUsecase()
	http.NewRequestHandler(mu)
}
