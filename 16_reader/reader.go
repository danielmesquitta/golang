//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package readers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readers() {
	r := bufio.NewReader(os.Stdin)

	commandsCount := map[string]int{}
	linesCount := 0

	for {
		line, err := r.ReadString('\n')
		line = strings.Trim(line, "\n")

		if err != nil {
			panic(err)
		}

		if line == "Q" || line == "q" {
			break
		}

		switch line {
		case "hello":
			fmt.Println("Hello!")
			commandsCount[line]++

		case "bye":
			fmt.Println("Bye!")
			commandsCount[line]++
		}

		linesCount++
	}

	fmt.Println("Commands:")
	fmt.Println(commandsCount)

	fmt.Println("Lines:")
	fmt.Println(linesCount)
}
