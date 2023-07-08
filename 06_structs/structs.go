//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package structs

import (
	"fmt"
)

type Rectangle struct {
	Length uint
	Width  uint
}

func (rectangle Rectangle) Area() uint {
	return rectangle.Length * rectangle.Width
}

func (rectangle Rectangle) Perimeter() uint {
	return 2*rectangle.Length + 2*rectangle.Width
}

func Structs() {
	rectangle := Rectangle{Length: 10, Width: 5}

	fmt.Println("Rectangle 1:", rectangle)
	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())

	rectangle.Length *= 2
	rectangle.Width *= 2

	fmt.Println("Rectangle 2:", rectangle)
	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())
}
