package variables

import "fmt"

func Variables() {
	// * Store your favorite color in a variable using the `var` keyword
	var favColor string = "blue"
	fmt.Println(favColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	var birthYear, age uint = 2002, 21
	fmt.Println(birthYear, age)

	//* Store your first & last initials in two variables using block assignment
	firstInitial, lastInitial := "A", "B"
	fmt.Println(firstInitial, lastInitial)

	//* Declare a variable for your age (in days),
	//  then assign it by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays uint = age * 365
	fmt.Println(ageInDays)
}
