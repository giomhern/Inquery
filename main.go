package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return "hello world"
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")

	if err != nil {
		log.Fatalf("err: %w", err)
	}

	defer f.Close()
	p := tea.NewProgram(model{}, tea.WithAltScreen())

}
