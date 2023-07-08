//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

//  * Print the result
//* Call every function at least once

package functions

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func sayHello() string {
	return "Hello"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func returnAnyNumber() int {
	return 1
}

// * Write a function that returns any two numbers
func returnAnyTwoNumbers() (int, int) {
	return 1, 2
}

// * Add three numbers together using any combination of the existing functions.
func addThreeNumbersUsingExistingFunctions() int {
	a, b := returnAnyTwoNumbers()
	c := returnAnyNumber()

	return addThreeNumbers(a, b, c)
}

func Functions() {
	greet("Daniel")

	fmt.Println(sayHello())

	fmt.Println(addThreeNumbers(1, 2, 3))

	fmt.Println(returnAnyNumber())

	fmt.Println(returnAnyTwoNumbers())

	fmt.Println(addThreeNumbersUsingExistingFunctions())
}
