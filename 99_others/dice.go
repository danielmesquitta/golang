//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package others

import (
	"fmt"
	"math/rand"
)

func Dice() {
	const initialRolls = 2
	const initialDices = 2
	const initialSides = 6

	sum := 0

	for rolls := 1; rolls <= initialRolls; rolls++ {
		fmt.Println("Roll " + fmt.Sprint(rolls) + ":")
		for dices := 1; dices <= initialDices; dices++ {
			diceRollResult := rand.Intn(initialSides) + 1

			fmt.Println("  Dice "+fmt.Sprint(dices)+":", diceRollResult)

			sum += diceRollResult
		}
	}

	fmt.Println("Sum:", sum)

	switch {
	case sum == 2 && initialDices == 2:
		fmt.Println("Snake eyes")
	case sum == 7:
		fmt.Println("Lucky 7")
	case sum%2 == 0:
		fmt.Println("Even")
	case sum%2 != 0:
		fmt.Println("Odd")
	default:
		fmt.Println("No special cases")
	}
}
