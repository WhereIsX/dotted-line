package main 

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/muesli/termenv" 
)


func main() {
	initCat := conciergeCat{
		services: [3]string{"create", "get", "delete"}
		currentService : -1
	}

	p := tea.NewProgram(initCat)
	if err := p.Start(); err != nil {
		fmt.Printf("You WAT: %v", err)
		os.Exit(1)
	}	
}

// To play nice with BubbleTea, we need:
// a model -- guestServices
// Init(), Update(), and View()

// the model 
type conciergeCat struct{
	services [3]string 
	currentService int // this points to which service we want 
}


// BubbleTea: Init
func (cc conciergeCat) Init() tea.Cmd {
	return nil
}

// BubbleTea: Update 
func (cc conciergeCat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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