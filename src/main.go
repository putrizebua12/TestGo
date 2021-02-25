package main

import (
	"apis/user_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NerRouter()

	router.HandleFunc("/api/user/findall", user_api.FindAll).Methods("GET")
	router.HandleFunc("/api/user/search/{keyword}", user_api.Search).Methods("GET")
	router.HandleFunc("/api/user/create", user_api.Create).Methods("POST")
	router.HandleFunc("/api/user/update", user_api.Update).Methods("PUT")
	router.HandleFunc("/api/user/delete/{id}", user_api.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
