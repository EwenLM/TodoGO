package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//Style
var (
	inputStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#c29af0"))
)

//Type

type model struct {
	input textinput.Model
	width int
	height int
}

type tickMsg struct{}

//Implementation

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.input.Width = msg.Width
		m.width = msg.Width
		m.height = msg.Height	
	}
	var inputCmd tea.Cmd
	m.input, inputCmd =m.input.Update(msg)
	return m, inputCmd
}
	

func (m model) View() string {
	return inputStyle.
	Width(m.width -2).
	Height(1).
	Render(
		m.input.View(),
	)
}

//Functions

func newModel() model {
	ti := textinput.New()
	ti.Placeholder = "Nouvelle tâche à faire"
	ti.Focus()
	return model{
		input: ti,
	}
}

func main() {
	m := newModel()
	tea.NewProgram(m).Run()

}
