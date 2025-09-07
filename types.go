package main

type DB struct {
	Version   int            `json:"version"`
	UpdatedAt string         `json:"updatedAt"`
	Skills    map[string]int `json:"skills"`
}

type StatusRow struct {
	Name                string
	Hours               float64
	Level               string
	HoursUntilNextLevel float64
	NextLabel           string
	PctToNext           float64
}

type Milestone struct {
	Hours float64
	Label string
}
