//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package iota

import "fmt"

type Operation byte

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

func (operation Operation) calculate(numbers ...int) int {
	result := numbers[0]

	switch operation {
	case Add:
		for _, number := range numbers[1:] {
			result += number
		}

	case Subtract:
		for _, number := range numbers[1:] {
			result -= number
		}

	case Multiply:
		for _, number := range numbers[1:] {
			result *= number
		}

	case Divide:
		for _, number := range numbers[1:] {
			result /= number
		}
	}

	return result
}

func Iota() {
	add, sub, mul, div := Add, Subtract, Multiply, Divide

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
