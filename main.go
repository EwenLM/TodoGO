package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//===Style===
var (
	inputStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#c29af0"))
)

//Type

type model struct {
	input textinput.Model
	list list.Model
	width int
	height int
}

type tickMsg struct{}

type todoItem struct {
	done bool
	name string
}



//====Implementation====

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.input.Value()
			m.input.SetValue("")
		}
	case tea.WindowSizeMsg:
		m.input.Width = msg.Width - 2
		m.list.SetSize(msg.Width, msg.Height - 3)
		m.width = msg.Width
		m.height = msg.Height	
	}
	var inputCmd tea.Cmd
	m.input, inputCmd =m.input.Update(msg)
	return m, inputCmd
}
	

func (m model) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		inputStyle.
		Width(m.width-2).
		Height(1).Render(
			m.input.View(),
		),
		m.list.View(),
	)
}

//===Functions===

func newModel() model {
	ti := textinput.New()
	ti.Placeholder = "Nouvelle tâche à faire"
	ti.Focus()
	li := list.New([]list.Item{
		todoItem{done : false, name : "Ménage Chambre"},
		todoItem{done : false, name : "Ménage SDB"},
		todoItem{done : true, name : "Ménage Salon"},
	}, list.NewDefaultDelegate(),0,0)
	return model{
		input: ti,
		list: li,
	}
}

func (t todoItem) FilterValue() string {
	return t.name
}

func (t todoItem) Description() string {
	return t.name
}
func (t todoItem) Title() string {
	return t.name
}

func main() {
	m := newModel()
	tea.NewProgram(m).Run()

}
