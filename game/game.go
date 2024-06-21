package game

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]*Player)
var broadcast = make(chan Message)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register new client
	clients[ws] = &Player{}
	log.Printf("New connection: %v", ws.RemoteAddr().String())
	notifyPlayerList()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			notifyPlayerList()
			break
		}
		handleMessage(ws, msg)
	}
}

func handleMessage(ws *websocket.Conn, msg Message) {
	switch msg.Type {
	case "join":
		clients[ws].Name = msg.Body
		clients[ws].Role = "Waiting"
		log.Printf("Player joined: %s", msg.Body)
		broadcast <- Message{Type: "info", Body: msg.Body + " has joined the game."}
		notifyPlayerList()
	case "start":
		assignRoles()
		for client := range clients {
			client.WriteJSON(Message{Type: "role", Body: clients[client].Role})
		}
		broadcast <- Message{Type: "info", Body: "Game has started!"}
	default:
		broadcast <- msg
	}
}

func notifyPlayerList() {
	var playerNames []string
	for _, player := range clients {
		if player.Name != "" {
			playerNames = append(playerNames, player.Name)
		}
	}
	log.Printf("Current players: %v", playerNames)
	playerList, _ := json.Marshal(playerNames)
	broadcast <- Message{Type: "player_list", Body: string(playerList)}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
				notifyPlayerList()
			}
		}
	}
}

func assignRoles() {
	var players []Player
	for _, player := range clients {
		players = append(players, *player)
	}

	numUndercover := len(players) / 4
	numCivilians := len(players) - numUndercover
	roles := AssignRoles(players, numUndercover, numCivilians)

	i := 0
	for client := range clients {
		clients[client].Role = roles[i].Role
		i++
	}
}
