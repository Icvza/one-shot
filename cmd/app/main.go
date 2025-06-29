package main

import (
	"fmt"
	"github.com/icvza/one-shot/internal/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
	prompt string
}

func initialModel() model {
	return model{
		prompt: "",
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("oneshot")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {

	s := "\nPress q to quit.\n"

	return s
}

func main() {
	ui.DisplayBanner()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
