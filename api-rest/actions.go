package main

import (
		"fmt"
        "net/http"
        "encoding/json"
        "github.com/gorilla/mux"
        
		)


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "El servidor está corriendo en http://localhost:8080")

}

func MovieList(w http.ResponseWriter, r *http.Request){
	movies := Movies{
	Movie{"sin limites",2013, "desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo Gas", 2005, "Juan ANtonio"},

	}

	//fmt.Fprintf(w, "listado de peliculas")
	json.NewEncoder(w).Encode(movies)

}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}