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
	focusedPromptColor = "132" // pastel pink?
	blurredPromptColor = "172" // rose gold ?
	focusedPrompt      = termenv.String("> ").Foreground(color("132")).String()
	blurredPrompt      = "  "
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

	focusedFormField := 0

	clientID := textinput.NewModel()
	clientID.Placeholder = "client ID"
	//clientID.CursorColor
	clientID.Prompt = focusedPrompt
	clientID.TextColor = focusedPromptColor
	clientID.Focus()

	authToken := textinput.NewModel()
	authToken.Placeholder = "auth token"
	//authToken.CursorColor
	authToken.Prompt = blurredPrompt
	authToken.TextColor = blurredPromptColor

	broadcasterUserID := textinput.NewModel()
	broadcasterUserID.Placeholder = "broadcaster user ID"
	//broadcasterUserID.CursorColor
	broadcasterUserID.Prompt = blurredPrompt
	broadcasterUserID.TextColor = blurredPromptColor

	initCat := conciergeCat{

		focusedFormField:  focusedFormField,
		clientID:          clientID,
		authToken:         authToken,
		broadcasterUserID: broadcasterUserID,

		view:    "signup",
		spinner: yanasSpinner,
	}

	return initCat
	// password.EchoMode = textinput.EchoPassword
	// password.EchoCharacter = '•'
}

// To play nice with BubbleTea, we need:
// a model -- conciergeCat
// Init(), Update(), and View()

// the model
type conciergeCat struct {
	// services [3]string
	// currentService int // this points to which service we want
	view string // options: "signup", "main", "create", "read", "delete", "delete_all"

	// signup Page Related
	userFinishedSignup bool
	focusedFormField   int
	clientID           textinput.Model
	authToken          textinput.Model
	broadcasterUserID  textinput.Model

	spinner spinner.Model
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
		case "ctrl+c", "q", "Q", "esc":
			return cc, tea.Quit
		}
	case spinner.TickMsg: // update spinner?
		var cmd tea.Cmd
		cc.spinner, cmd = cc.spinner.Update(msg)
		return cc, cmd
	}

	switch cc.view {
	case "signup":
		return UpdateSignupForm(cc, msg)
	}

	return cc, nil
}

func UpdateSignupForm(cc conciergeCat, msg tea.Msg) (tea.Model, tea.Cmd) {

	// up/down to change focused field
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			cc.focusedFormField--
		case "down":
			cc.focusedFormField++
		}
	}

	fields := []textinput.Model{
		cc.clientID,
		cc.authToken,
		cc.broadcasterUserID,
	}
	// keep focused field within bounds
	numFields := len(fields)
	if cc.focusedFormField < 0 {
		cc.focusedFormField = 0
	} else if cc.focusedFormField > numFields-1 {
		cc.focusedFormField = numFields - 1
	}

	for i := 0; i < numFields; i++ {
		if i == cc.focusedFormField {
			fields[i].Focus()
			fields[i].Prompt = focusedPrompt
		}
		fields[i].Blur()
		fields[i].Prompt = blurredPrompt
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
		display = cc.signupView()
	case "":
		display = defaultView
	default:
		display = defaultView
	}
	return display + cc.viewFooter()
}

func (cc conciergeCat) signupView() string {
	form := cc.clientID.View() + "\n" +
		cc.authToken.View() + "\n" +
		cc.broadcasterUserID.View() + "\n"

	return form
}

// the ever helperful footer :>
func (cc conciergeCat) viewFooter() string {
	return termenv.String("\n  ↑/↓: Navigate • q: Quit\n").Foreground(term.Color("241")).String()
}
