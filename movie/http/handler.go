package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"acb_task/movie/model"
	movieusecase "acb_task/movie/usecase"

	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

type MovieHandler struct {
	MovieUsecase movieusecase.MovieUsecase
}

func NewRequestHandler(mu movieusecase.MovieUsecase) {
	handler := &MovieHandler{
		MovieUsecase: mu,
	}
	myRouter := mux.NewRouter().StrictSlash(true)

	//Health Check API
	myRouter.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	myRouter.HandleFunc("/get_movies", handler.GetMovies).Methods("GET")
	myRouter.HandleFunc("/get_movie/{ID}", handler.GetMovieDetail).Methods("GET")
	myRouter.HandleFunc("/add_movie", handler.AddMovie).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}

func (m *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	resp, err := m.MovieUsecase.GetMovies(ctx)
	if err != nil {
		errMsg := errors.New("failed to get movies")
		log.Fatalln(errMsg)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)

}

func (m *MovieHandler) GetMovieDetail(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["ID"]
		ctx  = r.Context()
	)

	idNum, err := strconv.Atoi(id)
	if err != nil {
		errMsg := errors.New("invalid movie id")
		log.Fatalln(errMsg)
	}

	resp, err := m.MovieUsecase.GetMovieDetail(ctx, int64(idNum))
	if err != nil {
		errMsg := errors.New("failed to get movie data")
		log.Fatalln(errMsg)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)

}

func (m *MovieHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	var (
		request = new(model.AddMovieRequest)
		ctx     = r.Context()
	)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errMsg := errors.New("invalid request for adding movie")
		log.Fatalln(errMsg)
	}

	var ok bool
	if ok, err = isRequestValid(request); !ok {
		log.Fatalln("invalid request: ", err)
	}

	resp, err := m.MovieUsecase.AddMovie(ctx, *request)
	if err != nil {
		errMsg := errors.New("failed to add movie")
		log.Fatalln(errMsg)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)
}

func isRequestValid(m *model.AddMovieRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
