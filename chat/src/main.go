package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

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
	go gobot()
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
	if err != nil {
		log.Println(err)
	}
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

var addr = flag.String("addr", "localhost:8080", "http service address")

func gobot() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: ""}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
