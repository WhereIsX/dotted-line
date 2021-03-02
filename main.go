package main

import (
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	term               = termenv.ColorProfile()
	color              = termenv.ColorProfile().Color
	focusedPromptColor = "132" //
)

func main() {
	initCat := initialConciergeCat()
	p := tea.NewProgram(initCat)
	if err := p.Start(); err != nil {
		fmt.Printf("You WAT: %v", err)
		os.Exit(1)
	}
}

func initialConciergeCat() conciergeCat {
	yanasSpinner := spinner.NewModel()
	yanasSpinner.Spinner = spinner.MiniDot

	// clientID := textinput.NewModel()
	// clientID.Placeholder = "clientID"
	// //clientID.CursorColor
	// clientID.Prompt =
	// 	clientID.Focus()

	initCat := conciergeCat{

		view:    "signup",
		spinner: yanasSpinner,
	}

	return initCat

	// name := textinput.NewModel()
	// name.Placeholder = "Nickname"
	// name.Focus()
	// name.Prompt = focusedPrompt
	// name.TextColor = focusedTextColor
	// name.CharLimit = 32

	// email := textinput.NewModel()
	// email.Placeholder = "Email"
	// email.Prompt = blurredPrompt
	// email.CharLimit = 64

	// password := textinput.NewModel()
	// password.Placeholder = "Password"
	// password.Prompt = blurredPrompt
	// password.EchoMode = textinput.EchoPassword
	// password.EchoCharacter = '•'
	// password.CharLimit = 32

	// return model{0, name, email, password, blurredSubmitButton}

}

// To play nice with BubbleTea, we need:
// a model -- conciergeCat
// Init(), Update(), and View()

// the model
type conciergeCat struct {
	// services [3]string
	// currentService int // this points to which service we want
	view               string // options: "signup", "main", "create", "read", "delete", "delete_all"
	clientID           textinput.Model
	authToken          textinput.Model
	broadcasterUserID  textinput.Model
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

  PSST: you can dm yana github.com/whereisx and
  tell her her shit's broken`
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

func (cc conciergeCat) signupView() string {
	return ""
}

// the ever helperful footer :>
func (cc conciergeCat) viewFooter() string {
	return termenv.String("\n  ↑/↓: Navigate • q: Quit\n").Foreground(term.Color("241")).String()
}
