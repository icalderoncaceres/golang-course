package main

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"io"
	"time"
	"fmt"
)
func main(){
	//Rutas para la entidad campaign
	r := mux.NewRouter()
	r.HandleFunc("/campaigns/", handlerCampaignsGet).Methods("GET")
	r.HandleFunc("/campaign/{id}/", handlerCampaignGet).Methods("GET")
	r.HandleFunc("/campaigns/", handlerCampaignsPost).Methods("POST")
	r.HandleFunc("/campaigns/{id}/", handlerCampaignsPut).Methods("PUT")
	r.HandleFunc("/campaigns/{id}/", handlerCampaignsDelete).Methods("DELETE")

	//Arrancamos el servidor
	srv := &http.Server{
		Addr:           ":9080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
	fmt.Println("Is running on localhost:9080")
}



func handlerCampaignsGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code": 200,"message":"Campaign List"}`)
}

func handlerCampaignsPost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code":200,"message":"Campaign create"}`)
}

func handlerCampaignsPut(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code":200,"message":"Campaign update"`)
}

func handlerCampaignsDelete(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code":200,"message":"Campaign delete"`)
}

func handlerCampaignGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w, `{"code":200,"message":"Campaign show"`)	
}
