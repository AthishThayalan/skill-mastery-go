package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := loadFile()
	if err != nil {
		log.Fatal(err)
	}

	var db DB
	if err := json.Unmarshal(data, &db); err != nil {
		log.Fatal(err)
	}

	args := os.Args

	if len(args) > 1 && args[1] == "add" {
		skillName := args[2]
		duration, err := ParseDuration(args[3])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		AddSkillTime(skillName, duration, &db)
		if err := SaveFile(&db); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added %s to %s!\n\n", args[3], skillName)
		rows := computeRows(&db)
		printStatus(rows)
		return
	} else if len(args) == 2 && args[1] == "list" {
		rows := computeRows(&db)
		printStatus(rows)
	}
}
