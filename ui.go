package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	cReset  = "\033[0m"
	cBold   = "\033[1m"
	cDim    = "\033[2m"
	cBlue   = "\033[38;5;39m"
	cCyan   = "\033[36m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cMag    = "\033[35m"
	cGrey   = "\033[90m"
)

var skillColors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[91m", // Bright Red
	"\033[92m", // Bright Green
	"\033[93m", // Bright Yellow
	"\033[94m", // Bright Blue
	"\033[95m", // Bright Magenta
	"\033[96m", // Bright Cyan
}

func colorForLevel(level string) string {
	switch level {
	case "Getting Started":
		return cGrey
	case "Not bad":
		return cCyan
	case "Good":
		return cGreen
	case "Really good":
		return cBlue
	case "Amazing":
		return cMag
	case "World-class":
		return cYellow
	case "One of the best ever":
		return cBold + cYellow
	default:
		return cReset
	}
}

func bar(pct float64, width int) string {
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	filled := int((pct / 100.0) * float64(width))
	if filled > width {
		filled = width
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

var ansiRE = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRE.ReplaceAllString(s, "")
}

func padANSI(s string, width int) string {
	vis := len([]rune(stripANSI(s)))
	if vis >= width {
		return s
	}
	return s + strings.Repeat(" ", width-vis)
}

func colorForSkill(name string) string {
	// simple hash of skill name
	sum := 0
	for _, ch := range name {
		sum += int(ch)
	}
	return skillColors[sum%len(skillColors)]
}

func printStatus(rows []StatusRow) {
	if len(rows) == 0 {
		fmt.Println("No skills tracked yet. Add your first one.")
		return
	}

	// Column widths
	const (
		colSkill = 18
		colHours = 8
		colLevel = 22
		colNextH = 14
		colPct   = 9
	)

	title := cBold + cBlue + "Skill Progress (to next milestone)" + cReset
	sep := strings.Repeat("─", colSkill+2+colHours+2+colLevel+2+colNextH+2+colPct+8)
	fmt.Println(title)
	fmt.Println(sep)

	// Header
	hSkill := padANSI(cDim+"Skill"+cReset, colSkill)
	hHours := padANSI(cDim+"Hours"+cReset, colHours)
	hLevel := padANSI(cDim+"Level"+cReset, colLevel)
	hNextH := padANSI(cDim+"Hours→Next"+cReset, colNextH)
	hPct := padANSI(cDim+"%→Next"+cReset, colPct)
	fmt.Printf("%s  %s  %s  %s  %s\n", hSkill, hHours, hLevel, hNextH, hPct)
	fmt.Println(sep)

	for _, r := range rows {
		nextH := "—"
		if r.HoursUntilNextLevel > 0 {
			nextH = fmt.Sprintf("%.2f", r.HoursUntilNextLevel)
		}

		levelColor := colorForLevel(r.Level)
		level := levelColor + r.Level + cReset

		coloredSkill := colorForSkill(r.Name) + r.Name + cReset
		skillCell := padANSI(coloredSkill, colSkill)
		hoursCell := padANSI(fmt.Sprintf("%.2f", r.Hours), colHours)
		levelCell := padANSI(level, colLevel)
		nextHCell := padANSI(nextH, colNextH)
		pctCell := padANSI(fmt.Sprintf("%6.2f%%", r.PctToNext), colPct)

		fmt.Printf("%s  %s  %s  %s  %s\n", skillCell, hoursCell, levelCell, nextHCell, pctCell)
	}

	fmt.Println(sep)
}
