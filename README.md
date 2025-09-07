# Skill Mastery Tracker

A simple terminal-based app inspired by Malcolm Gladwell’s **10,000-hour rule**.  
Track your practice time across different skills and see your progress toward mastery.

## Motivation

I wanted a lightweight tool to **quantify deliberate practice** without bloated apps.  
This tracker uses a JSON file for persistence and maps your hours to milestones:

- **100h** → Not bad
- **1,000h** → Good
- **2,000h** → Really good
- **5,000h** → Amazing
- **10,000h** → World-class
- **20,000h** → One of the best ever

## Usage

### Show progress

```bash
go run . list
```

### Add Time

```bash
go run . add <activity> <time>
```

### Examples

```bash
go run . add Go 1h
go run . add Guitar 30m
```
