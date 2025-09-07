package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DB struct {
	Version   int            `json:"version"`
	UpdatedAt string         `json:"updatedAt"`
	Skills    map[string]int `json:"skills"`
}

type StatusRow struct {
	Name      string
	Hours     float64
	Level     string
	NextLabel string
	PctTo10k  float64
}

type Milestone struct {
	Hours float64
	Label string
}

var Milestones = []Milestone{
	{Hours: 100, Label: "Not bad"},
	{Hours: 1000, Label: "Good"},
	{Hours: 2000, Label: "Really good"},
	{Hours: 5000, Label: "Amazing"},
	{Hours: 10000, Label: "World-class"},
	{Hours: 20000, Label: "One of the best ever"},
}

func computeRows(db *DB) []StatusRow {
	var rows []StatusRow

	for skillName, mins := range db.Skills {
		name := skillName
		hours := float64(mins / 60)
		nextLabel := nextMilestone(hours)
		level := computeLevel(hours)
		pct := pctTo(hours)
		rows = append(rows, StatusRow{Name: name, Hours: hours, Level: level, NextLabel: nextLabel, PctTo10k: pct})
	}

	return rows
}

func loadFile() ([]byte, error) {
	path := filepath.Join("skills", "skills.json")
	return os.ReadFile(path)
}

func main() {
	data, err := loadFile()

	if err != nil {
		panic(err)
	}

	var db DB
	if err := json.Unmarshal(data, &db); err != nil {
		panic(err)
	}

	rows := computeRows(&db)
	printStatus(rows)

}

func printStatus(rows []StatusRow) {
	if len(rows) == 0 {
		fmt.Println("No skills tracked yet. Add your first one.")
		return
	}

	// Header
	fmt.Printf("%-20s %-10s %-18s %-12s %-8s\n",
		"Skill", "Hours", "Level", "→ Next", "% of 10k")
	fmt.Println(strings.Repeat("-", 75))

	// Rows
	for _, r := range rows {
		next := "—"
		if r.NextLabel != "" {
			next = r.NextLabel
		}
		fmt.Printf("%-20s %-10.2f %-18s %-12s %-8.2f%%\n",
			r.Name, r.Hours, r.Level, next, r.PctTo10k)
	}
}

func nextMilestone(hours float64) string {
	for _, milestone := range Milestones {
		if hours < milestone.Hours {
			return milestone.Label
		}
	}
	return ""
}

func pctTo(hours float64) float64 {
	return (hours / 10000) * 100
}

func computeLevel(hours float64) string {
	if hours < 100 {
		return "Getting Started"
	}
	for i, _ := range Milestones {
		if i == 0 {
			continue
		}
		if hours < Milestones[i].Hours {
			return Milestones[i-1].Label
		}
	}
	return ""
}
