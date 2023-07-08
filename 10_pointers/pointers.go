//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package pointers

import "fmt"

const (
	Active   = true
	Inactive = false
)

type Item struct {
	Name        string
	SecurityTag bool
}

func activateTag(item *Item) {
	item.SecurityTag = Active
}

func deactivateTag(item *Item) {
	item.SecurityTag = Inactive
}

func checkout(items *[4]Item) {
	for index := range *items {
		itemPtr := &(*items)[index]
		deactivateTag(itemPtr)
	}
}

func printItems(items *[4]Item) {
	fmt.Println("Items:")

	for _, item := range *items {
		fmt.Println(item.Name, ":", item.SecurityTag)
	}

	fmt.Println()
}

func Pointers() {
	var items [4]Item

	for index := range items {
		itemPtr := &items[index]
		itemPtr.Name = "Item " + fmt.Sprint(index+1)
		activateTag(itemPtr)
	}

	printItems(&items)

	deactivateTag(&items[1])

	printItems(&items)

	checkout(&items)

	printItems(&items)
}
