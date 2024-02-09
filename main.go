package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

type model struct {
	index       int
	questions   []string
	width       int
	height      int
	answerField textinput.Model
	styles      *Styles
}

type Question struct {
	question string
	answer   string
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

func New(questions []string) *model {
	styles := DefaultStyles()
	answerField := textinput.New()
	answerField.Placeholder = "Type your answer here"
	answerField.Focus()
	return &model{questions: questions, answerField: answerField, styles: styles}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg: // on start up ==> only once it shows up
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg: // a specific key message
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.Next()
			m.answerField.SetValue(("done!"))
			return m, nil
		}
	}
	newValue, newCmd := m.answerField.Update(msg)
	m.answerField = newValue
	return m, newCmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, m.questions[m.index], m.styles.InputField.Render(m.answerField.View())))
}

func (m *model) Next() {
	if m.index < (len(m.questions) - 1) {
		m.index++
	} else {
		m.index = 0
	}
}

func main() {
	questions := []string{"what is your name?", "what is your favorite editor?", "what is your quote?"}
	m := New(questions)
	f, err := tea.LogToFile("debug.log", "debug")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
