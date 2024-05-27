package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

var testNames = []string{"Paavo", "Gigachad", "Harold", "Taavi", "MogBot"}

type wsMsg struct {
	// senderID string
	name    string
	message string
}

type model struct {
	viewport       viewport.Model
	messages       []string
	textarea       textarea.Model
	senderStyle    lipgloss.Style
	recipientStyle lipgloss.Style
	ws             *websocket.Conn
	err            error
}

type (
	errMsg error
)

func InitialModel(ws *websocket.Conn) model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "| "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room! Type a message and press enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:       ta,
		messages:       []string{},
		viewport:       vp,
		senderStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("36")),
		recipientStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		ws:             ws,
		err:            nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			_, err := m.ws.Write([]byte(m.textarea.Value()))
			if err != nil {
				log.Println("Error sending message: ", err)
			}
			// 	m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())

			// For testing
			// m.messages = append(m.messages, m.recipientStyle.Render("MogBot: ")+"Based")

			// m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}
	case wsMsg:
		m.messages = append(m.messages, m.senderStyle.Render(msg.name+": ")+msg.message)
		m.viewport.SetContent(strings.Join(m.messages, "\n"))

		m.viewport.GotoBottom()

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)

}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	// Websocket connection
	url := os.Getenv("SERVERURL")
	origin := os.Getenv("CLIENT_ORIGIN")

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Print the logo
	asciiArt, err := os.ReadFile("./assets/ascii_art.txt")
	if err != nil {
		log.Println("Error loading the ascii art: ", err)
	}

	fmt.Println(string(asciiArt))

	p := tea.NewProgram(InitialModel(ws))

	// Message handler
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Println("Error reading the message: ", err)
				return
			}
			p.Send(wsMsg{
				name:    testNames[rand.IntN(4)],
				message: msg})
		}
	}()

	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
