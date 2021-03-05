// powered by zongzi -- yana
// derived from github.com/charmbracelet/bubbletea/blob/master/examples/textinputs
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
)

const focusedTextColor = "205"

var (
	color               = te.ColorProfile().Color
	focusedPrompt       = te.String("> ").Foreground(color("205")).String()
	blurredPrompt       = "> "
	focusedSubmitButton = "[ " + te.String("Submit").Foreground(color("205")).String() + " ]"
	blurredSubmitButton = "[ " + te.String("Submit").Foreground(color("240")).String() + " ]"
)

func main() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

type model struct {
	focusedField  int
	nameInput     textinput.Model
	emailInput    textinput.Model
	passwordInput textinput.Model
	submitButton  string

	view string
}

func initialModel() model {
	name := textinput.NewModel()
	name.Placeholder = "Nickname"
	name.Focus()
	name.Prompt = focusedPrompt
	name.TextColor = focusedTextColor
	name.CharLimit = 32

	email := textinput.NewModel()
	email.Placeholder = "Email"
	email.Prompt = blurredPrompt
	email.CharLimit = 64

	password := textinput.NewModel()
	password.Placeholder = "Password"
	password.Prompt = blurredPrompt
	password.EchoMode = textinput.EchoPassword
	password.EchoCharacter = '•'
	password.CharLimit = 32

	return model{
		focusedField:  0,
		nameInput:     name,
		emailInput:    email,
		passwordInput: password,
		submitButton:  focusedSubmitButton,
		view:          "",
	}

}
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "esc", "q", "Q":
			return m, tea.Quit

		// Cycle between inputs
		case "tab", "shift+tab", "enter", "up", "down":

			inputs := []textinput.Model{
				m.nameInput,
				m.emailInput,
				m.passwordInput,
			}

			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusedField == len(inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusedField--
			} else {
				m.focusedField++
			}

			// implenting wraparound for menu
			if m.focusedField > len(inputs) {
				m.focusedField = 0
			} else if m.focusedField < 0 {
				m.focusedField = len(inputs)
			}

			for i := 0; i <= len(inputs)-1; i++ {
				if i == m.focusedField {
					// Set focused state
					inputs[i].Focus()
					inputs[i].Prompt = focusedPrompt
					inputs[i].TextColor = focusedTextColor
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].Prompt = blurredPrompt
				inputs[i].TextColor = ""
			}

			m.nameInput = inputs[0]
			m.emailInput = inputs[1]
			m.passwordInput = inputs[2]

			if m.focusedField == len(inputs) {
				m.submitButton = focusedSubmitButton
			} else {
				m.submitButton = blurredSubmitButton
			}

			return m, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateInputs(msg, m)
	return m, cmd
}

// Pass messages and models through to text input components. Only text inputs
// with Focus() set will respond, so it's safe to simply update all of them
// here without any further logic.
func updateInputs(msg tea.Msg, m model) (model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.nameInput, cmd = m.nameInput.Update(msg)
	cmds = append(cmds, cmd)

	m.emailInput, cmd = m.emailInput.Update(msg)
	cmds = append(cmds, cmd)

	m.passwordInput, cmd = m.passwordInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	s := "\n"

	inputs := []string{
		m.nameInput.View(),
		m.emailInput.View(),
		m.passwordInput.View(),
	}

	// logic for adding newlines between inputs
	for i := 0; i < len(inputs); i++ {
		s += inputs[i]
		if i < len(inputs)-1 {
			s += "\n"
		}
	}

	s += "\n\n" + m.submitButton + "\n"
	return s
}
