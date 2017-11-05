package main 

import (
	//	"fmt"
        "net/http"
        "log"
       // "encoding/json"
       // "github.com/gorilla/mux"
        
		)

func main(){

    /*router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)

    router.HandleFunc("/peliculas", MovieList)
    router.HandleFunc("/pelicula/{id}", MovieShow)*/

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		fmt.Fprintf(w, "Hola mundo desde mi servidor Web con GO")
	})*/

	router := NewRouter()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

	


}



