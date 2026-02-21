package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ===Style===
var (
	inputStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#c29af0"))
)

//Type

type model struct {
	input  textinput.Model
	list   list.Model
	width  int
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
			todo:= todoItem {name: m.input.Value()}
			m.input.SetValue("")
			cmd := m.list.InsertItem(0,todo)
			return m, cmd
		case tea.KeyDown:
			if m.input.Focused(){
				m.input.Blur()
				return m, nil
			}
		case tea.KeyUp:
			if m.input.Focused()== false && m.list.Cursor()== 0{
				m.input.Focus()
				return m,nil
			}
		}
	case tea.WindowSizeMsg:
		m.input.Width = msg.Width - 2
		m.list.SetSize(msg.Width, msg.Height-3)
		m.width = msg.Width
		m.height = msg.Height
	}
	var inputCmd, listCmd tea.Cmd
	if m.input.Focused() {
		m.input, inputCmd = m.input.Update(msg)
	} else {
		m.list, listCmd = m.list.Update(msg)
	}

	return m, tea.Batch(inputCmd, listCmd)
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
		todoItem{done: false, name: "Ménage Chambre"},
		todoItem{done: false, name: "Ménage SDB"},
		todoItem{done: true, name: "Ménage Salon"},
	}, list.NewDefaultDelegate(), 0, 0)
	li.SetShowStatusBar(false)
	li.DisableQuitKeybindings()
	return model{
		input: ti,
		list:  li,
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
	tea.NewProgram(m, tea.WithAltScreen()).Run()

}
