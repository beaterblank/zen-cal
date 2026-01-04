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
	titleStyle, headerStyle,
		footerStyle, weekNumStyle,
		weekdayStyle, weekendStyle,
		todayStyle := getStyles()

	firstWeekDay, week, lastDay := getMonthInfo(c.selMonth, c.selYear)

	var cal strings.Builder

	headers := [8]string{"#", "Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	for _, header := range headers {
		cal.WriteString(headerStyle.Render(header))
	}
	cal.WriteByte('\n')

	cal.WriteString(weekNumStyle.Render(fmt.Sprintf("%d", week)))

	currWeekDay := int(firstWeekDay)
	for i := 0; i < currWeekDay; i++ {
		// Render an empty string with the cell style to maintain 4-char width
		cal.WriteString(weekdayStyle.Render(""))
	}

	for d := 1; d <= lastDay; d++ {
		if currWeekDay%7 == 0 && d != 1 {
			cal.WriteByte('\n')
			week++
			cal.WriteString(weekNumStyle.Render(fmt.Sprintf("%d", week)))
		}

		dayStr := fmt.Sprintf("%d", d)
		var style lipgloss.Style

		if d == c.currDay && c.currYear == c.selYear && c.selMonth == c.currMonth {
			style = todayStyle
		} else if currWeekDay%7 == 0 || currWeekDay%7 == 6 {
			style = weekendStyle
		} else {
			style = weekdayStyle
		}

		cal.WriteString(style.Render(dayStr))
		currWeekDay++
	}

	calStr := cal.String()
	calWidth := lipgloss.Width(calStr)

	title := titleStyle.
		Width(calWidth).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("%s %d", c.selMonth, c.selYear))

	footer := footerStyle.
		Width(calWidth).
		Align(lipgloss.Center).
		Render("\n⇄ month  ⇅ year  ↵ reset")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		calStr,
		footer,
	)
}
