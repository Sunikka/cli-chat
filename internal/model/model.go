package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model for frontend menu

type model struct {
	choice string
}

// func initialModel() model {
// 	return model{
// 		choices: []string{"Start chatting", "Join a Group", "Exit"},
// 	}
// }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl + c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// Figure out later
		case "enter", " ":
		}
	}
	return m, nil
}
func (m model) View() string {

}

func choicesView(m model) string
