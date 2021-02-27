package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"os"
)

var (
	term  = termenv.ColorProfile()
	color = termenv.ColorProfile().Color
)

// spinner.Line,
// spinner.Dot,
// spinner.MiniDot,
// spinner.Jump,
// spinner.Pulse,
// spinner.Points,
// spinner.Globe,
// spinner.Moon,
// spinner.Monkey,

func main() {
	yanasSpinner := spinner.NewModel()
	yanasSpinner.Spinner = spinner.MiniDot
	initCat := conciergeCat{
		clientID:          "test",
		authToken:         "test",
		broadcasterUserID: "more tests",
		// services: [3]string{"create", "get", "delete"},
		view:    "signup",
		spinner: yanasSpinner,
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
	spinner            spinner.Model
}

// BubbleTea: Init
func (cc conciergeCat) Init() tea.Cmd {
	return spinner.Tick
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
	case spinner.TickMsg: // update spinner?
		var cmd tea.Cmd
		cc.spinner, cmd = cc.spinner.Update(msg)
		return cc, cmd
	}
	return cc, nil
}

// BubbleTea: View
func (cc conciergeCat) View() string {
	var display string
	defaultStringLiteral := `
	welp. we couldn't find anything nice to show 
	so here's a spinner for your troubles %s

	PSST: you can dm yana github.com/whereisx and tell her her shit's broken`
	spinner := termenv.String(cc.spinner.View()).Foreground(color("205")).String()
	defaultView := fmt.Sprintf(defaultStringLiteral, spinner)

	switch cc.view {
	case "signup":
		// display = signupView()
		display = defaultView
	case "":
		display = defaultView
	default:
		display = defaultView
	}
	return display + cc.viewFooter()
}

// the ever helperful footer :>
func (cc conciergeCat) viewFooter() string {
	return termenv.String("\n  ↑/↓: Navigate • q: Quit\n").Foreground(term.Color("241")).String()
}
