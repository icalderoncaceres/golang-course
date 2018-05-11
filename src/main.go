package main

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"time"
	"io"
	"fmt"

)
func main(){
	//Rutas para la entidad campaign
	r := mux.NewRouter()
	r.HandleFunc("/repositories/{search}", handlerRepositoriesGet).Methods("GET") //Buscar una lista de repositorios dependiendiendo de la palabra clave
	r.HandleFunc("/repository/{key}", handlerRepositoryGet).Methods("GET") //Traer un repositorio en concreto

	//Arrancamos el servidor
	srv := &http.Server{
        Addr:         ":9080",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
	}
	log.Fatal(srv.ListenAndServe())
	fmt.Print("Server is running on port 9080")
}



func handlerRepositoriesGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")	
	io.WriteString(w, `{"code": 200,"message":"Repositories List with key=`)
}

func handlerRepositoryGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code":200,"message":"Repository show"`)	
}