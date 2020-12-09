package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"url_shortener/config"
	"url_shortener/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.Load()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", controllers.CreateShortURL).Methods("POST")
	router.HandleFunc("/{url}", controllers.GetURL).Methods("GET")

	corsOps := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", strconv.Itoa(config.PORT)),
		Handler:      corsOps.Handler(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()

}
