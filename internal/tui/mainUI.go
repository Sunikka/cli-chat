package tui

import tea "github.com/charmbracelet/bubbletea"

var p *tea.Program

type sessionState int

const (
	loginView sessionState = iota
	chatView
)

type MainModel struct {
	state sessionState
	login tea.Model
	chat  tea.Model
}

func (m MainModel) Init() tea.Model {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmdi

	switch msg = msg.(type) {
		case 
	}
}

func (m MainModel) View() string {
	switch m.state {
	case chatView:
		return m.chat.View()

	default:
		return m.login.View()
	}

}
