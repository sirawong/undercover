package main

import (
	"log"
	"net/http"
	"undercover/game"
)

func main() {
	http.HandleFunc("/ws", game.HandleConnections)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server started on :8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
