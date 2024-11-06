package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		} else if command == "update" {
			if len(notes) == 0 {
				fmt.Println("[Error] There is nothing to update")
				continue
			}

			if len(input) == 1 {
				fmt.Println("[Error] Missing position argument")
			} else if len(input) < 3 {
				fmt.Println("[Error] Missing note argument")
			} else {
				id, err := strconv.Atoi(input[1])
				if err != nil {
					fmt.Printf("[Error] Invalid position: %s\n\n", input[1])
					continue
				}

				update(&notes, uint(id), strings.Join(input[2:], " "))
			}
		} else if command == "delete" {

			if len(notes) == 0 {
				fmt.Println("[Error] There is nothing to delete")
				continue
			}

			if len(input) < 2 {
				fmt.Println("[Error] Missing position argument")
			} else {
				id, err := strconv.Atoi(input[1])
				if err != nil {
					fmt.Printf("[Error] Invalid position: %s\n\n", input[1])
					continue
				}
				delete(&notes, uint(id))
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

func update(notes *[]string, id uint, note string) {

	if id > uint(len(*notes)) {
		fmt.Printf(getLimitStr(id))
		return
	}

	(*notes)[id-1] = note
	fmt.Printf("[OK] The note at position %d was successfully updated\n", id)
}

func delete(notes *[]string, id uint) {
	if id == 0 || int(id) > len(*notes) {
		fmt.Printf(getLimitStr(id))
		return
	}

	index := id - 1
	*notes = append((*notes)[:index], (*notes)[index+1:]...)

	fmt.Printf("[OK] The note at position %d was successfully deleted\n", id)
}

func getLimitStr(id uint) string {
	limit := []int{}

	for i := 1; i <= maxNote; i++ {
		limit = append(limit, i)
	}

	limitStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(limit)), ", "), "[]")
	return fmt.Sprintf("[Error] Position %d is out of the boundaries [%v]\n", id, limitStr)
}
