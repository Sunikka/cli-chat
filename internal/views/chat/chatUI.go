package chatUI

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/net/websocket"
	"os"
)

var testNames = []string{"Paavo", "Gigachad", "Harold", "Taavi", "MogBot"}

type wsMsg struct {
	// senderID string
	name    string
	message string
}

type Model struct {
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

func InitialModel(ws *websocket.Conn) Model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "| "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(100, 25)
	vp.SetContent(fmt.Sprintf("Welcome to the chat room!\nType a message and press enter to send."))

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return Model{
		textarea:       ta,
		messages:       []string{},
		viewport:       vp,
		senderStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("36")),
		recipientStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		ws:             ws,
		err:            nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())

			// For testing
			// m.messages = append(m.messages, m.recipientStyle.Render("MogBot: ")+"Based")

			m.viewport.SetContent(strings.Join(m.messages, "\n"))
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

func (m Model) View() string {
	// Print the logo
	asciiArt, err := os.ReadFile("assets/ascii_art.txt")
	if err != nil {
		log.Println("Error loading the ascii art: ", err)
	}

	return fmt.Sprintf(
		"%s\n%s\n\n%s",
		asciiArt,
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
