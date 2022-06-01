package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	movieusecase "acb_task/movie/usecase"
)

type MovieHandler struct {
	MovieUsecase movieusecase.MovieUsecase
}

func NewRequestHandler(mu movieusecase.MovieUsecase) {
	handler := &MovieHandler{
		MovieUsecase: mu,
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/get_movie/{ID}", handler.GetMovie).Methods("GET")
	//myRouter.HandleFunc("/movie_detail/{imdbID}", handler.GetMovieDetail).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("ACB_Movie REST API listening on port", port)

	err := http.ListenAndServe(":8080", myRouter)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}

}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["ID"]
	)

	resp, err := m.MovieUsecase.GetMovieDetail(id)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)

}

// func (m *MovieHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		vars       = mux.Vars(r)
// 		imdbID     = vars["imdbID"]
// 		reqBody, _ = ioutil.ReadAll(r.Body)
// 		req        model.Request
// 	)

// 	json.Unmarshal(reqBody, &req)

// 	resp, err := m.MovieUsecase.GetMovieDetail(imdbID)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	w.Write(resp)

// }