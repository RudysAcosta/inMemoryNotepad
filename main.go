package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxNote int

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var notes []string

	fmt.Printf("Enter the maximum number of notes:")
	fmt.Scanln(&maxNote)

	for {
		fmt.Print("Enter a command and data:")
		scanner.Scan()

		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		if command == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		} else if command == "create" {
			if strings.Trim(data, " ") == "" {
				fmt.Println("[Error] Missing note argument")
			} else {
				create(&notes, data)
			}
		} else if command == "clear" {
			notes = notes[:0]
			fmt.Println("[OK] All notes were successfully deleted")
		} else if command == "list" {
			if len(notes) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				list(&notes)
			}
		} else {
			fmt.Println("[Error] Unknown command")
		}

		fmt.Println()

	}
}

func create(notes *[]string, note string) {
	if len(*notes) < maxNote {
		*notes = append(*notes, note)
		fmt.Println("[OK] The note was successfully created")
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}

func list(notes *[]string) {
	for index, note := range *notes {
		fmt.Printf("[Info] %d: %s\n", (index + 1), note)
	}
}
