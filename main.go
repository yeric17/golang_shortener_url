package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"url_shortener/config"
	"url_shortener/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.CreateShortURL).Methods("POST")
	router.HandleFunc("/{url}", controllers.GetURL).Methods("GET")

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", strconv.Itoa(config.PORT)),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()

}
