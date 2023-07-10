//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func sumFileLines(fileName string) (int, error) {
	sum := 0

	readFile, err := os.Open("20_goroutines/" + fileName)

	if err != nil {
		return 0, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	for _, line := range fileLines {
		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Error:", err)

			continue
		}

		sum += number
	}

	return sum, nil
}

func Channels() {
	start := time.Now()

	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	sumsChannel := make(chan int)

	for _, file := range files {
		fileName := file

		go func() {
			sum, err := sumFileLines(fileName)

			if err != nil {
				fmt.Println("Error:", err)
			}

			sumsChannel <- sum
		}()
	}

	grandTotal := 0

	for range files {
		grandTotal += <-sumsChannel
	}

	fmt.Println("Grand Total:", grandTotal)

	elapsed := time.Since(start)
	fmt.Println("Time Elapsed:", elapsed)
}

func WaitGroups() {
	start := time.Now()

	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	grandTotal := 0

	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)

		fileName := file

		go func() {
			sum, err := sumFileLines(fileName)

			if err != nil {
				fmt.Println("Error:", err)
			}

			grandTotal += sum

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Grand Total:", grandTotal)

	elapsed := time.Since(start)
	fmt.Println("Time Elapsed:", elapsed)
}
