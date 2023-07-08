//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package slices

import "fmt"

type Part string

func printAssemblyLine(assemblyLine []Part) {
	fmt.Println("Assembly line:")

	for _, part := range assemblyLine {
		fmt.Println(part)
	}

	fmt.Println()
}

func Slices() {
	assemblyLine := []Part{"Part 1", "Part 2", "Part 3"}

	printAssemblyLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Part 4", "Part 5")

	printAssemblyLine(assemblyLine)

	assemblyLine = assemblyLine[3:]

	printAssemblyLine(assemblyLine)
}
