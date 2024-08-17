package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sunikka/termitalk/internal/routes"
	wsHandler "github.com/Sunikka/termitalk/internal/routes"
	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	serverPort := os.Getenv("SERVERPORT")
	loginPort := os.Getenv("AUTH_PORT")
	go startLoginService(loginPort)

	startMainService(serverPort)
}

func startLoginService(port string) {
	http.HandleFunc("/login", routes.HandleLogin)

	log.Println("Login Service listening on port", port)
	http.ListenAndServe(port, nil)
}

func startMainService(port string) {
	server := wsHandler.NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleConn))

	log.Println("Server listening on port", port)
	http.ListenAndServe(port, nil)
}
