package main

import (
	"fmt"
	"net/http"

	wsHandler "github.com/Sunikka/termitalk/internal/handlers"
	"golang.org/x/net/websocket"
)

const PORT string = ":3000"

func main() {
	server := wsHandler.NewServer()

	http.Handle("/ws", websocket.Handler(server.HandleConn))

	fmt.Println("Server listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
