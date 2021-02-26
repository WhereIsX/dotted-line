package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"os"
)

var term = termenv.ColorProfile()

func main() {
	initCat := conciergeCat{
		clientID:          "test",
		authToken:         "test",
		broadcasterUserID: "more tests",
		// services: [3]string{"create", "get", "delete"},
		view: "signup",
	}

	p := tea.NewProgram(initCat)
	if err := p.Start(); err != nil {
		fmt.Printf("You WAT: %v", err)
		os.Exit(1)
	}
}

// To play nice with BubbleTea, we need:
// a model -- conciergeCat
// Init(), Update(), and View()

// the model
type conciergeCat struct {
	// services [3]string
	// currentService int // this points to which service we want
	view               string // options: "main", "create", "read", "delete", "delete_all"
	clientID           string
	authToken          string
	broadcasterUserID  string
	userFinishedSignup bool
}

// BubbleTea: Init
func (cc conciergeCat) Init() tea.Cmd {
	return nil
}

// BubbleTea: Update
func (cc conciergeCat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // keyboard input?!
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return cc, tea.Quit
		case "up", "k":

		case "down", "j":

		case "enter", " ":

		}
	}
	return cc, nil
}

// BubbleTea: View
func (cc conciergeCat) View() string {

	display := "test test"
	return display + cc.viewFooter()
}

// the ever helperful footer :>
func (cc conciergeCat) viewFooter() string {
	return termenv.String("\n  ↑/↓: Navigate • q: Quit\n").Foreground(term.Color("241")).String()
}
