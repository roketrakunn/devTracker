package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Entry     string    `json:"entry"`
	Tag       string    `json:"tag,omitempty"`
}

type XpStats struct {
	Total int `json:"total"`
}

const logFile = "log.json"
const xpFile = "xp.json"

var validXpGain = map[string]int{
"Learned Go":                 120,
	"Learned Rust":               120,
	"Did easy leetcode":          20,
	"Did medium leetcode":        40,
	"Did hard leetcode":          70,
	"Learned a new vim motion/Trick": 40,
}

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	tag := addCmd.String("tag", "", "Optional tag for the entry")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'today', 'xp', or 'progress' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() < 1 {
			fmt.Println("Usage: devtrack add \"your message\" --tag=optionalTag")
			os.Exit(1)
		}
		entry := strings.Join(addCmd.Args(), " ")
		saveLog(entry, *tag)

	case "today":
		showToday()

	case "xp":
		stats := loadXp()
		fmt.Printf("Total XP: %d\n", stats.Total)

	case "progress":
		showProgress()
	case "streak": 
		showStreak()

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func saveLog(message, tag string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Entry:     message,
		Tag:       tag, //could be how you feel hahaha... 
	}

	var logs []LogEntry
	data, _ := os.ReadFile(logFile)
	if len(data) > 0 {
		json.Unmarshal(data, &logs)
	}

	logs = append(logs, entry)

	updated, _ := json.MarshalIndent(logs, "", "  ")
	os.WriteFile(logFile, updated, 0644)

	fmt.Println("Logged:", message)
	rewardXp(entry)
}

func rewardXp(entry LogEntry) {
	for key, xp := range validXpGain {
		if strings.Contains(strings.ToLower(entry.Entry), strings.ToLower(key)) {
			fmt.Printf("Gained %d XP for: %q\n", xp, entry.Entry)

			stats := loadXp()
			stats.Total += xp
			saveXp(stats)

			return
		}
	}
	fmt.Println("Definitely a waste of time!")
}

func loadXp() XpStats {
	var stats XpStats
	data, err := os.ReadFile(xpFile)
	if err != nil {
		return stats
	}
	json.Unmarshal(data, &stats)
	return stats
}

func saveXp(stats XpStats) {
	data, _ := json.MarshalIndent(stats, "", "  ")
	os.WriteFile(xpFile, data, 0644)
}

func showToday() {
	var logs []LogEntry
	data, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("No logs found yet.")
		return
	}
	json.Unmarshal(data, &logs)

	today := time.Now().Format("2006-01-02")
	found := false
	for _, log := range logs {
		if log.Timestamp.Format("2006-01-02") == today {
			fmt.Printf("[%s] %s %s\n", log.Timestamp.Format("15:04"),
				log.Entry, emojiTag(log.Tag))
			found = true
		}
	}
	if !found {
		fmt.Println("No entries for today yet.")
	}
}

func emojiTag(tag string) string {
	if tag == "" {
		return ""
	}
	return fmt.Sprintf("%s", tag)
}

func showProgress() {
	var logs []LogEntry
	data, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("No logs found yet.")
		return
	}
	json.Unmarshal(data, &logs)

	progress := make(map[string]int)

	for _, log := range logs {
		date := log.Timestamp.Format("Monday") // or "2006-01-02"
		for key, xp := range validXpGain {
			if strings.Contains(strings.ToLower(log.Entry), strings.ToLower(key)) {
				progress[date] += xp
				break
			}
		}
	}

	if len(progress) == 0 {
		fmt.Println("No XP progress to show yet.")
		return
	}

	fmt.Println("Weekly XP Progress:")
	for day, xp := range progress {
		fmt.Printf("%s: %d XP\n", day, xp)
	}

	TotalXp := 0//For overall Xp gained to this time

	for _,xp := range progress{ 
		TotalXp +=xp
	}

	fmt.Println("________________________________")	
	fmt.Printf("You have a total of %dXP", TotalXp)

}


func  showStreak(){ 

	var logs []LogEntry
	data ,err := os.ReadFile(logFile)

	if err !=nil{ 
		fmt.Println("No logs availbe yet")
		return
	}


	json.Unmarshal(data , &logs)

	dateMap := make(map[string]bool)

	for _,log := range logs{ 

		day := log.Timestamp.Format("2006-01-02")

		dateMap[day] = true 
	}

	streak := 0 


	for i := 0 ; ;i++{ 

		day := time.Now().AddDate(0 , 0 ,-i).Format("2006-01-02")

		if dateMap[day]{ 

			streak++
		} else { 
			break
		}
	}

	if streak ==0{ 
		fmt.Println("No streak yet?,You better start working man!")

	}else { 
		fmt.Printf("Current streak %d day(s)\n", streak)
	}
}


