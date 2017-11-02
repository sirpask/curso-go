package main 

import (
		"fmt"
        "net/http"
        "log"
        "github.com/gorilla/mux"
		)

func main(){

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)

    router.HandleFunc("/contacto", Contact)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		fmt.Fprintf(w, "Hola mundo desde mi servidor Web con GO")
	})*/

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

	


}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "El servidor está corriendo en http://localhost:8080")

}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Esta es la pagina de Contacto")

}