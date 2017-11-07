package main

import (
		"fmt"
        "net/http"
        "encoding/json"
        "github.com/gorilla/mux"
        "log"
		)

var movies = Movies{
	Movie{"sin limites",2013, "desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo Gas", 2005, "Juan ANtonio"},
	}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "El servidor está corriendo en http://localhost:8080")

}

func MovieList(w http.ResponseWriter, r *http.Request){
	/*movies := Movies{
	Movie{"sin limites",2013, "desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo Gas", 2005, "Juan ANtonio"},

	}*/

	//fmt.Fprintf(w, "listado de peliculas")
	json.NewEncoder(w).Encode(movies)

}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

func MovieAdd(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
    err := decoder.Decode(&movie_data)

    if (err != nil){

    	panic(err)
    }

    defer r.Body.Close()

    log.Println(movie_data)
    movies = append(movies, movie_data)

	}