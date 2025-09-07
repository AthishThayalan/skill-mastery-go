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

func printStatus(rows []StatusRow) {
	if len(rows) == 0 {
		fmt.Println("No skills tracked yet. Add your first one.")
		return
	}

	const (
		colSkill = 18
		colHours = 8
		colLevel = 22
		colNext  = 10
		colPct   = 8
	)

	title := cBold + cBlue + "Skill Progress (toward 10k)" + cReset
	sep := strings.Repeat("─", colSkill+2+colHours+2+colLevel+2+colNext+2+colPct+10)
	fmt.Println(title)
	fmt.Println(sep)

	hSkill := padANSI(cDim+"Skill"+cReset, colSkill)
	hHours := padANSI(cDim+"Hours"+cReset, colHours)
	hLevel := padANSI(cDim+"Level"+cReset, colLevel)
	hNext := padANSI(cDim+"→ Next"+cReset, colNext)
	hPct := cDim + "% of 10k" + cReset
	fmt.Printf("%s  %s  %s  %s  %s\n", hSkill, hHours, hLevel, hNext, hPct)
	fmt.Println(sep)

	for _, r := range rows {
		next := "—"
		if r.HoursUntilNextLevel > 0 {
			next = fmt.Sprintf("%.2fh", r.HoursUntilNextLevel)
		}

		levelColor := colorForLevel(r.Level)
		level := levelColor + r.Level + cReset

		skillCell := padANSI(r.Name, colSkill)
		hoursCell := padANSI(fmt.Sprintf("%.2f", r.Hours), colHours)
		levelCell := padANSI(level, colLevel)
		nextCell := padANSI(next, colNext)
		pctCell := fmt.Sprintf("%6.2f%%", r.PctTo10k)

		fmt.Printf("%s  %s  %s  %s  %s\n", skillCell, hoursCell, levelCell, nextCell, pctCell)
	}

	fmt.Println(sep)
}
