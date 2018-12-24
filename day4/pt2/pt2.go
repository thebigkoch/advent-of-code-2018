package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func getGuardsList(messages map[string]string) map[int]int {
	result := make(map[int]int)
	var currentGuardId int

	for _, message := range messages {
		if strings.Contains(message, "begins shift") {
			fmt.Sscanf(message, "Guard #%d begins shift", &currentGuardId)
			result[currentGuardId] += 1
		}
	}

	return result
}

// Inputs:
//    kDateTimes: The sorted date-times of messages
//    messages: The messages that were found
//    guardId: The ID of the guard to track
// Output: A mapping.
//    Keys = Number of minutes after midnight
//    Values = Number of days that a guard slept during that minute
func calculateSleepingMinutes(kDateTimes []string, messages map[string]string, guardId int) map[int]int {
	result := make(map[int]int)
	var currentGuardId int = 0
	var sleepHour, sleepMinute int

	for _, strDateTime := range kDateTimes {
		message := messages[strDateTime]
		if strings.Contains(message, "begins shift") {
			fmt.Sscanf(message, "Guard #%d begins shift", &currentGuardId)
		} else if currentGuardId == guardId {
			if strings.Contains(message, "falls asleep") {
				sleepHour, _ = strconv.Atoi(strDateTime[11:13])
				sleepMinute, _ = strconv.Atoi(strDateTime[14:16])

				// Ignore sleeping before midnight
				if sleepHour == 23 {
					sleepMinute = 0
				}
			} else if strings.Contains(message, "wakes up") {
				wakeHour, _ := strconv.Atoi(strDateTime[11:13])
				wakeMinute, _ := strconv.Atoi(strDateTime[14:16])

				// Ignore waking before midnight
				if wakeHour == 0 {
					for i := sleepMinute; i < wakeMinute; i++ {
						result[i] = result[i] + 1
					}
				}
			}
		}
	}

	return result
}

func main() {
	// Get the current directory.
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Open the input file.
	filePath := filepath.Join(currentDir, "input.txt")
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	var kDateTimes []string
	messages := make(map[string]string)
	scanner := bufio.NewScanner(f)

	// Read all lines into temporary variables.  Read the keys into temporary variable kDates.
	for scanner.Scan() {
		currentLine := scanner.Text()
		strDateTime := currentLine[1:17]
		strMsg := currentLine[19:]
		messages[strDateTime] = strMsg
		kDateTimes = append(kDateTimes, strDateTime)
	}

	// Sort date-times.
	sort.Strings(kDateTimes)

	// Get the list of guards
	guardsList := getGuardsList(messages)

	// Find the number of days that the given guard was asleep each minute.  Get the highest value.
	var selectedGuardId int = 0
	var selectedMinute int = 0
	var selectedMinutesSpentAsleep int = 0
	for guardId, _ := range guardsList {
		mapSleepingMinutes := calculateSleepingMinutes(kDateTimes, messages, guardId)
		for k, v := range mapSleepingMinutes {
			if v > selectedMinutesSpentAsleep {
				selectedGuardId = guardId
				selectedMinute = k
				selectedMinutesSpentAsleep = v
			}
		}
	}
	fmt.Printf("GuardID: %d, Minute: %d\n", selectedGuardId, selectedMinute)
}
