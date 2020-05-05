package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/movie", MovieListHandler)
	r.HandleFunc("/movie/{id:[0-9]+}", MovieHandler)

	http.ListenAndServe(":8080", r)
}

func MovieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(MovieList())
	return
}

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(MovieList()[id])
	return
}

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieUrl string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}

func MovieList() []*Movie {
	return []*Movie{
		&Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		&Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		&Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}
}

// &Movie{3, "Железный человек", "/static/posters/ironman.jpg"},
// &Movie{4, "Холодное сердце", "/static/posters/cold.jpg"},
// &Movie{5, "Эверест", "/static/posters/everest.jpg"},
// &Movie{6, "Пираты карибского моря", "/static/posters/pirates.jpg"},
// &Movie{7, "Оно", "/static/posters/it.jpg"},
// &Movie{8, "Однажды в голливуде", "/static/posters/hw.jpg"},
// &Movie{9, "Звездные войны", "/static/posters/starwars.jpg"},
