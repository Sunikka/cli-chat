package main

import (
	mainUI "github.com/Sunikka/termitalk/internal/views/main"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
	"log"
	"math/rand/v2"
	"os"
)

var testNames = []string{"Paavo", "Gigachad", "Harold", "Taavi", "MogBot"}

type wsMsg struct {
	// senderID string
	name    string
	message string
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	// Websocket connection
	url := os.Getenv("SERVERURL")
	origin := os.Getenv("CLIENT_ORIGIN")

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	app := tea.NewProgram(mainUI.NewMainModel(ws), tea.WithAltScreen())

	// Message handler
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Println("Error reading the message: ", err)
				return
			}
			app.Send(wsMsg{
				name:    testNames[rand.IntN(4)],
				message: msg})
		}
	}()

	_, err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
