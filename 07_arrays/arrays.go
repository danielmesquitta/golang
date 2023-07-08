//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package arrays

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func getShoppingListStatus(shoppingList []Product) (int, float64) {
	length, totalPrice := 0, 0.0

	for _, product := range shoppingList {
		if product.Name != "" {
			totalPrice += product.Price
			length++
		}
	}

	return length, totalPrice
}

func printShoppingListStatus(shoppingList []Product) {
	length, totalPrice := getShoppingListStatus(shoppingList)

	fmt.Println("Last item:", shoppingList[length-1].Name)

	fmt.Println("Cart length:", length)

	fmt.Println("Total price:", totalPrice)
}

func Array() {
	shoppingList := [4]Product{
		{Name: "Milk", Price: 1.99},
		{Name: "Eggs", Price: 2.99},
		{Name: "Bread", Price: 0.99},
	}

	printShoppingListStatus(shoppingList[:])

	shoppingList[3] = Product{Name: "Butter", Price: 1.99}

	printShoppingListStatus(shoppingList[:])

}
