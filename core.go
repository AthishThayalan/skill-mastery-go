package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func computeRows(db *DB) []StatusRow {
	var rows []StatusRow

	for skillName, mins := range db.Skills {
		hours := float64(mins) / 60.0

		nextHours, nextLabel := nextMilestone(hours)
		var hrsToNext float64
		if nextLabel != "" && nextHours > hours {
			hrsToNext = nextHours - hours
		}

		level := computeLevel(hours)
		pct := pctTo(hours)

		rows = append(rows, StatusRow{
			Name:                skillName,
			Hours:               hours,
			Level:               level,
			HoursUntilNextLevel: hrsToNext,
			NextLabel:           nextLabel,
			PctTo10k:            pct,
		})
	}

	return rows
}

func nextMilestone(hours float64) (float64, string) {
	for _, milestone := range Milestones {
		if hours < milestone.Hours {
			return milestone.Hours, milestone.Label
		}
	}
	return 0, ""
}

func pctTo(hours float64) float64 {
	p := (hours / 10000.0) * 100.0
	if p < 0 {
		return 0
	}
	if p > 100 {
		return 100
	}
	return p
}

func computeLevel(hours float64) string {
	if hours < 100 {
		return "Getting Started"
	}
	for i := 0; i < len(Milestones); i++ {
		if hours < Milestones[i].Hours {
			return Milestones[i-1].Label
		}
	}

	return Milestones[len(Milestones)-1].Label
}

func ParseDuration(s string) (int, error) {
	if strings.HasSuffix(s, "m") {
		minutes, err := strconv.Atoi(strings.TrimSuffix(s, "m"))
		return minutes, err
	} else if strings.HasSuffix(s, "h") {
		hours, err := strconv.Atoi(strings.TrimSuffix(s, "h"))
		return hours * 60, err
	}
	return 0, fmt.Errorf("invalid duration format (use 1h or 15m)")
}

func AddSkillTime(skill string, mins int, db *DB) {
	if db.Skills == nil {
		db.Skills = make(map[string]int)
	}
	db.Skills[skill] += mins
}

func SaveFile(db *DB) error {
	path := filepath.Join("skills", "skills.json")
	data, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
