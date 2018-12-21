package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	chatbot bot.Bot
	clients map[*Client]bool

	broadcastmsg chan *ClientMessage
	register     chan *Client
	unregister   chan *Client
}

func NewHub(chatbot bot.Bot) *Hub {

	return &Hub{
		chatbot:      chatbot,
		broadcastmsg: make(chan *ClientMessage), register: make(chan *Client), unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) ChatBot() bot.Bot {

	return h.chatbot
}

func (h *Hub) SendMessage(client *Client, message []byte) {

	client.send
}
