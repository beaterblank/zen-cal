package main

import (
	"time"

	lipgloss "github.com/charmbracelet/lipgloss"
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

func getMonthInfo(month time.Month, year int) (time.Weekday, int, int) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstWeekDay := firstDay.Weekday() // Sunday = 0
	lastDay := firstDay.AddDate(0, 1, -1).Day()
	jan1 := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	week := (firstDay.YearDay()+int(jan1.Weekday())-1)/7 + 1
	return firstWeekDay, week, lastDay
}

func getPallete() (lipgloss.Color, lipgloss.Color, lipgloss.Color, lipgloss.Color) {
	subtle := lipgloss.Color("241")    // Dim gray
	accent := lipgloss.Color("63")     // Soft slate blue
	highlight := lipgloss.Color("231") // Near white
	weekend := lipgloss.Color("210")   // Soft salmon/muted red
	return subtle, accent, highlight, weekend
}

func getStyles() (lipgloss.Style, lipgloss.Style, lipgloss.Style, lipgloss.Style, lipgloss.Style, lipgloss.Style, lipgloss.Style) {
	subtle, accent, highlight, weekend := getPallete()
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(highlight)

	footerStyle := lipgloss.NewStyle()

	cellStyle := lipgloss.NewStyle().
		Width(4). // Increased width for breathing room
		Align(lipgloss.Center)

	headerStyle := cellStyle.
		Foreground(subtle).
		Italic(true)

	weekNumStyle := cellStyle.
		Foreground(lipgloss.Color("239")).
		Italic(true)

	weekdayStyle := cellStyle.Foreground(lipgloss.Color("250"))
	weekendStyle := cellStyle.Foreground(weekend)

	todayStyle := cellStyle.
		Foreground(highlight).
		Background(accent).
		Bold(true)

	return titleStyle, headerStyle, footerStyle, weekNumStyle, weekdayStyle, weekendStyle, todayStyle
}
