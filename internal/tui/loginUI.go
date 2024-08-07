package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	viewport  viewport.Model
	userField textarea.Model
	pwField   textarea.Model
	err       error
}
