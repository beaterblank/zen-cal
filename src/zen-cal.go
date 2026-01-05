package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(newCalendarPage())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
