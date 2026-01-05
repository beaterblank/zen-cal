package main

import (
	"fmt"
	"strings"

	lipgloss "github.com/charmbracelet/lipgloss"
)

// VIEW
func (c calendarPage) View() string {
	return buildCal(c)
}

func buildCal(c calendarPage) string {
	styles := c.styles

	firstWeekDay, week, lastDay := getMonthInfo(c.selMonth, c.selYear)

	var cal strings.Builder

	// Header row: week numbers + weekdays
	headers := [8]string{"#", "Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	for _, header := range headers {
		cal.WriteString(styles.headerStyle.Render(header))
	}
	cal.WriteByte('\n')

	// Start with first week number
	cal.WriteString(styles.weekNumStyle.Render(fmt.Sprintf("%d", week)))

	currWeekDay := int(firstWeekDay)
	// Fill empty days before the first of the month
	for i := 0; i < currWeekDay; i++ {
		cal.WriteString(styles.weekdayStyle.Render("    ")) // 4-char width placeholder
	}

	// Render days of the month
	for d := 1; d <= lastDay; d++ {
		if currWeekDay%7 == 0 && d != 1 {
			cal.WriteByte('\n')
			week++
			cal.WriteString(styles.weekNumStyle.Render(fmt.Sprintf("%d", week)))
		}

		dayStr := fmt.Sprintf("%d", d)
		var style lipgloss.Style

		switch {
		case d == c.currDay && c.currYear == c.selYear && c.selMonth == c.currMonth:
			style = styles.todayStyle
		case currWeekDay%7 == 0 || currWeekDay%7 == 6:
			style = styles.weekendStyle
		default:
			style = styles.weekdayStyle
		}

		cal.WriteString(style.Render(dayStr))
		currWeekDay++
	}

	// Calculate calendar width for centering title and footer
	calStr := cal.String()
	calWidth := lipgloss.Width(calStr)

	title := styles.titleStyle.
		Width(calWidth).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("%s %d", c.selMonth, c.selYear))

	footer := styles.footerStyle.
		Width(calWidth).
		Align(lipgloss.Center).
		Render("⇄ month  ⇅ year  ↵ reset")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		calStr,
		footer,
	)
}
