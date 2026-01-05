package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
func (c calendarPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return c, tea.Quit

		case "j", "down": // increase year
			if c.selYear < 9999 {
				c.selYear++
			}

		case "k", "up": // decrease year
			if c.selYear > 1 {
				c.selYear--
			}

		case "l", "right": // increase month
			if c.selMonth == time.December {
				c.selMonth = time.January
				c.selYear++
			} else {
				c.selMonth++
			}

		case "h", "left": // decrease month
			if c.selMonth == time.January {
				c.selMonth = time.December
				c.selYear--
			} else {
				c.selMonth--
			}

		case "r", "enter": // get new curr time and reset today
			year, month, day := time.Now().Date()
			c.currDay = day
			c.currMonth = month
			c.currYear = year
			c.selMonth = month
			c.selYear = year
		}
	}
	return c, nil
}
