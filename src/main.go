package main

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"os"

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
        Addr:         ":9080",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
	}
	log.Fatal(srv.ListenAndServe())
}



func handlerCampaignsGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	os.WriteString(w, `{"code": 200,"message":"Campaign List"}`)
}

func handlerCampaignsPost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	os.WriteString(w, `{"code":200,"message":"Campaign create"}`)
}

func handlerCampaignsPut(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	os.WriteString(w, `{"code":200,"message":"Campaign update"`)
}

func handlerCampaignsDelete(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	os.WriteString(w, `{"code":200,"message":"Campaign delete"`)
}

func handlerCampaignGet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(htt.StatusOK)
	w.Header().Set("Content-Type","application/json")
	os.WriteString(w, `{"code":200,"message":"Campaign show"`)	
}
