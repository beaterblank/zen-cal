package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	styles    calstyle
}

type calstyle struct {
	titleStyle   lipgloss.Style
	headerStyle  lipgloss.Style
	footerStyle  lipgloss.Style
	weekNumStyle lipgloss.Style
	weekdayStyle lipgloss.Style
	weekendStyle lipgloss.Style
	todayStyle   lipgloss.Style
}

func newCalendarPage() calendarPage {
	year, month, day := time.Now().Date()
	return calendarPage{
		currDay:   day,
		currMonth: month,
		currYear:  year,
		selMonth:  month,
		selYear:   year,
		styles:    getStyles(),
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

func getPalette() (primary, secondary, text, alert lipgloss.Color) {
	primary = lipgloss.Color("63")    // Accent / highlight background
	secondary = lipgloss.Color("241") // Subtle / dim text
	text = lipgloss.Color("231")      // Main text / foreground
	alert = lipgloss.Color("210")     // Special highlight (weekend / alert)
	hexColor := regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return // Return defaults if home directory cannot be determined
	}
	
	configPath := filepath.Join(homeDir, ".config", "zen-cal", "zen-cal.conf")
	file, err := os.Open(configPath)
	if err != nil {
		return // Return defaults if config file doesn't exist
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		if !hexColor.MatchString(val) {
			continue // skip invalid hex codes
		}
		switch key {
		case "primary":
			primary = lipgloss.Color(val)
		case "secondary":
			secondary = lipgloss.Color(val)
		case "text":
			text = lipgloss.Color(val)
		case "alert":
			alert = lipgloss.Color(val)
		}
	}
	
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		// Return defaults if there was an error reading the file
		return
	}
	
	return
}

func getStyles() calstyle {
	primary, secondary, text, alert := getPalette()

	// Base cell style
	cellBase := lipgloss.NewStyle().
		Width(4).
		Align(lipgloss.Center)

	return calstyle{
		// Title / main headings
		titleStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(text),

		// Footer
		footerStyle: lipgloss.NewStyle().
			Foreground(text),

		// Header / index / row numbers
		headerStyle: cellBase.
			Foreground(secondary).
			Italic(true),

		weekNumStyle: cellBase.
			Foreground(secondary).
			Italic(true),

		// Weekday cells
		weekdayStyle: cellBase.
			Foreground(text),

		// Weekend / alert cells
		weekendStyle: cellBase.
			Foreground(alert),

		// Current day / active selection
		todayStyle: cellBase.
			Foreground(text).
			Background(primary).
			Bold(true),
	}
}
