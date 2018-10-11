package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) //lista połączonych klientów
var broadcast = make(chan Message)           //kanał broadcast

var upgrader = websocket.Upgrader{} //zmienia http na websocked

//Definicja typu Message
type Message struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	//Tworzenie serwera
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	//Zdefiniowanie trasy websocket
	http.HandleFunc("/ws", handleConnectios)

	//Rozpoczęcie nasłuchiwania przychodzących wiadomości
	go handleMessages()

	//Uruchomienie serwera i logowanie błędów
	log.Println("Serwer uruchomiony na porcie :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnectios(w http.ResponseWriter, r *http.Request) {
	//Upgrade get to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Zamknięcie połączenia
	defer ws.Close()

	//Dodanie nowych klientów do mapy
	clients[ws] = true

	for {
		var msg Message

		//Odczytanie wiadomości JSON i zmapowanie na obiekt Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		//Wysłanie wiadomości na broadcast
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		//Pobranie kolejnej wiadomości z kanału
		msg := <-broadcast

		//Rozesłanie wiadomości
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
