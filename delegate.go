package main

import (
	"io"

	"github.com/charmbracelet/bubbles/list"
)

type todoDelegate struct {
	focused bool
}

// Pointer
func newTodoDelegate() *todoDelegate {
	return &todoDelegate{}
}

func (t todoDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	panic("")
}

func (t todoDelegate) Spacing() int {
	panic("")
}

func (t todoDelegate) Update(msg tea.Msg, m *list.Model) tea.cmd {
	panic("")
}

