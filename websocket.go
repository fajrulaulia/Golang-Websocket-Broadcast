package beechatt_socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Payload struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Message  string `json:"message"`
	Callback bool   `json:"callback"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *Payload)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func writer(payload *Payload) {
	broadcast <- payload
}

func PayloadCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == http.MethodOptions {
		return
	}
	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}

	defer r.Body.Close()
	go writer(&payload)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"message\":\"Success\"}")))

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"message\":\"Welcome to home\"}")))

}

func PayloadWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true
}

func Echo() {
	for {
		val := <-broadcast
		latlong, _ := json.Marshal(val)
		for client := range clients {
			log.Println(val.From + " SEND TO " + val.To + " With Message : " + val.Message)
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
				return
			}
		}
	}
}
