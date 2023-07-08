//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package others

import (
	"fmt"
	"time"
)

type Member struct {
	Name string
}

type CheckOut struct {
	At     time.Time
	Member *Member
}

type CheckIn struct {
	At     time.Time
	Member *Member
}

type Book struct {
	Name        string
	CheckedOuts []CheckOut
	CheckedIns  []CheckIn
}

type Library struct {
	Books   *[]Book
	Members *[]Member
}

func printLibrary(library *Library) {
	fmt.Println("Books:")
	for _, book := range *library.Books {
		fmt.Println(book.Name)

		fmt.Println("CheckedIns:")
		for _, checkOut := range book.CheckedIns {
			fmt.Println(checkOut.At, checkOut.Member.Name)
		}
		fmt.Println()

		fmt.Println("CheckedOuts:")
		for _, checkOut := range book.CheckedOuts {
			fmt.Println(checkOut.At, checkOut.Member.Name)
		}

		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Members:")
	for _, member := range *library.Members {
		fmt.Println(member.Name)
	}

	fmt.Println()
}

func checkOut(book *Book, member *Member) {
	book.CheckedOuts = append(book.CheckedOuts, CheckOut{
		At:     time.Now(),
		Member: member,
	})
}

func checkIn(book *Book, member *Member) {
	book.CheckedIns = append(book.CheckedIns, CheckIn{
		At:     time.Now(),
		Member: member,
	})
}

func LibraryExercise() {
	books := []Book{
		{
			Name: "Book 1",
		},
		{
			Name: "Book 2",
		},
		{
			Name: "Book 3",
		},
		{
			Name: "Book 4",
		},
	}

	members := []Member{
		{
			Name: "Member 1",
		},
		{
			Name: "Member 2",
		},
		{
			Name: "Member 3",
		},
	}

	library := Library{
		Books:   &books,
		Members: &members,
	}

	printLibrary(&library)

	checkIn(&books[0], &members[0])

	printLibrary(&library)

	checkOut(&books[0], &members[0])

	printLibrary(&library)
}
