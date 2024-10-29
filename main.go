package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data:")
		scanner.Scan()

		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		if command == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		}

		fmt.Println(command, data)
		fmt.Println()
	}
}
