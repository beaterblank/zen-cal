package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	table "github.com/charmbracelet/lipgloss/table"
)

// MODEL DATA
type calendarPage struct {
	currDay   int
	currMonth time.Month
	currYear  int
	selMonth  time.Month
	selYear   int
}

func newCalendarPage() calendarPage {
	year, month, day := time.Now().Date()
	return calendarPage{
		currDay:   day,
		currMonth: month,
		currYear:  year,
		selMonth:  month,
		selYear:   year,
	}
}

// Init
func (c calendarPage) Init() tea.Cmd { return tea.ClearScreen }

// VIEW
func (c calendarPage) View() string {
	return buildCal(c)
}

// UPDATE
func (c calendarPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return c, tea.Quit

		case "j": // increase year
			c.selYear++

		case "k": // decrease year
			c.selYear--

		case "l": // increase month
			if c.selMonth == time.December {
				c.selMonth = time.January
				c.selYear++
			} else {
				c.selMonth++
			}

		case "h": // decrease month
			if c.selMonth == time.January {
				c.selMonth = time.December
				c.selYear--
			} else {
				c.selMonth--
			}

		case "r": // get new curr time and reset today
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

func getMonthInfo(month time.Month, year int) (time.Weekday, int) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstWeekDay := firstDay.Weekday() // Sunday = 0
	lastDay := firstDay.AddDate(0, 1, -1).Day()
	return firstWeekDay, lastDay
}

func buildCal(c calendarPage) string {
	// --- Palette ---
	subtle := lipgloss.Color("241")    // Dim gray
	accent := lipgloss.Color("63")     // Soft slate blue
	highlight := lipgloss.Color("231") // Near white
	weekend := lipgloss.Color("210")   // Soft salmon/muted red

	// --- Styles ---
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(highlight)

	cellStyle := lipgloss.NewStyle().
		Width(4). // Increased width for breathing room
		Align(lipgloss.Center)

	// Header (S M T W...)
	headerStyle := cellStyle.
		Foreground(subtle).
		Italic(true)

	// Week numbers (the left column)
	weekNumStyle := cellStyle.
		Foreground(lipgloss.Color("239")).
		Italic(true)

	// Standard days
	weekdayStyle := cellStyle.Foreground(lipgloss.Color("250"))
	weekendStyle := cellStyle.Foreground(weekend)

	// Modern "Today" Style: Subtle but distinct
	todayStyle := cellStyle.
		Foreground(highlight).
		Background(accent).
		Bold(true)

	// --- Logic ---
	firstWeekDay, lastDay := getMonthInfo(c.selMonth, c.selYear)
	var rows [][]string
	var currentRow []string

	// Using ISO week logic or simple increment
	week := 1
	currentRow = append(currentRow, fmt.Sprintf("%d", week))

	for i := 0; i < int(firstWeekDay); i++ {
		currentRow = append(currentRow, "")
	}

	currWeekDay := int(firstWeekDay)
	for d := 1; d <= lastDay; d++ {
		if currWeekDay%7 == 0 && d != 1 {
			rows = append(rows, currentRow)
			week++
			currentRow = []string{fmt.Sprintf("%d", week)}
		}
		currentRow = append(currentRow, fmt.Sprintf("%d", d))
		currWeekDay++
	}

	for len(currentRow) < 8 {
		currentRow = append(currentRow, "")
	}
	rows = append(rows, currentRow)

	// --- Table Rendering ---
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers("W#", "S", "M", "T", "W", "Th", "F", "Sa").
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}

			// Week Number Column
			if col == 0 {
				return weekNumStyle
			}

			dayStr := rows[row][col]
			if dayStr == "" {
				return cellStyle
			}

			// Check if Today
			isToday := c.selMonth == c.currMonth &&
				c.selYear == c.currYear &&
				dayStr == fmt.Sprintf("%d", c.currDay)

			if isToday {
				return todayStyle
			}

			// Weekend Coloring
			if col == 1 || col == 7 {
				return weekendStyle
			}

			return weekdayStyle
		})

	title := titleStyle.Render(fmt.Sprintf("%s %d", c.selMonth, c.selYear))

	return lipgloss.JoinVertical(lipgloss.Left, title, t.String())
}

func main() {
	p := tea.NewProgram(newCalendarPage())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
