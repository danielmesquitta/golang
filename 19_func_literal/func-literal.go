//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package funcLiterals

import "unicode"

func iterateLines(lines []string, callback func(string)) {
	for _, line := range lines {
		callback(line)
	}
}

func FuncLiterals() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	var lettersCount, digitsCount, spacesCount, punctuationsCount int

	iterateCharacters := func(line string) {
		for _, char := range line {
			switch {
			case unicode.IsLetter(char):
				lettersCount++
			case unicode.IsDigit(char):
				digitsCount++
			case unicode.IsSpace(char):
				spacesCount++
			case unicode.IsPunct(char):
				punctuationsCount++
			}
		}
	}

	iterateLines(lines, iterateCharacters)

	println("Letters:", lettersCount)
	println("Digits:", digitsCount)
	println("Spaces:", spacesCount)
	println("Punctuations:", punctuationsCount)
}
