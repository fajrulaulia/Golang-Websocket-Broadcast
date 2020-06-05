package main

import (
	c "beechatt-socket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("Web server run on PORT 8084")

	router := mux.NewRouter()
	router.HandleFunc("/payload", c.Index).Methods("GET")
	router.HandleFunc("/payload", c.PayloadCreate).Methods("POST", "OPTIONS")
	go c.Echo()

	cors := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: true,
	})

	router.HandleFunc("/ws", c.PayloadWs)

	log.Fatal(http.ListenAndServe(":8084", cors.Handler(router)))

}
