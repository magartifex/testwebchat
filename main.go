package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"testwebchat/client"
	"testwebchat/websocket"
	"testwebchat/websocket/messages"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/ws", websocket.Handler)
	http.HandleFunc("/client", client.Handler)

	go messages.GarbageСollectorWorker()

	// Порт я бы вынес в конфиг
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", client.Data.Port), nil))
}
