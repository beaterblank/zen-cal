package main

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
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

func IsDarkModeGTK() bool {
	cmd := exec.Command(
		"gsettings",
		"get",
		"org.gnome.desktop.interface",
		"color-scheme",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return false
	}

	// 'prefer-dark' or 'default'
	return strings.Contains(out.String(), "dark")
}

func getMonthInfo(month time.Month, year int) (time.Weekday, int, int) {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstWeekDay := firstDay.Weekday() // Sunday = 0
	lastDay := firstDay.AddDate(0, 1, -1).Day()
	jan1 := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	week := (firstDay.YearDay()+int(jan1.Weekday())-1)/7 + 1
	return firstWeekDay, week, lastDay
}

func getPalette() (today, headings, text, weekends lipgloss.Color) {
	today = lipgloss.Color("#f38ba8")    // Accent / highlight background
	headings = lipgloss.Color("#cba6f7") // Subtle / dim text
	text = lipgloss.Color("#cdd6f4")     // Main text / foreground
	weekends = lipgloss.Color("#f9e2af") // Special highlight (weekend / weekends)
	hexColor := regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return // Return defaults if home directory cannot be determined
	}
	// defaults to light mode
	var confFileName string
	if IsDarkModeGTK() {
		confFileName = "zen-cal-dark.conf"
	} else {
		confFileName = "zen-cal-light.conf"
	}
	configPath := filepath.Join(homeDir, ".config", "zen-cal", confFileName)
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
		case "today":
			today = lipgloss.Color(val)
		case "headings":
			headings = lipgloss.Color(val)
		case "text":
			text = lipgloss.Color(val)
		case "weekends":
			weekends = lipgloss.Color(val)
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
	today, headings, text, weekends := getPalette()

	// Base cell style
	cellBase := lipgloss.NewStyle().
		Width(4).
		Align(lipgloss.Center)

	return calstyle{
		// text / main headings
		titleStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(text),

		// Footer
		footerStyle: lipgloss.NewStyle().
			Foreground(text),

		// Header / index / row numbers
		headerStyle: cellBase.
			Foreground(headings).
			Italic(true),

		weekNumStyle: cellBase.
			Foreground(headings).
			Italic(true),

		// Weekday cells
		weekdayStyle: cellBase.
			Foreground(text),

		// Weekend / weekends cells
		weekendStyle: cellBase.
			Foreground(weekends),

		// Current day / active selection
		todayStyle: cellBase.
			Foreground(text).
			Background(today).
			Bold(true),
	}
}
