//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package errorsPkg

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}

type InvalidTimeError struct {
	receivedStr string
}

func (i *InvalidTimeError) Error() string {
	return fmt.Sprintf("invalid time string: got %s, want HH:MM:SS", i.receivedStr)
}

func ParseTime(str string) (Time, error) {
	timeSlice := strings.Split(str, ":")

	if len(timeSlice) != 3 {
		return Time{}, &InvalidTimeError{str}
	}

	hours, err := strconv.Atoi(timeSlice[0])
	if err != nil {
		return Time{}, err
	}

	minutes, err := strconv.Atoi(timeSlice[1])
	if err != nil {
		return Time{}, err
	}

	seconds, err := strconv.Atoi(timeSlice[2])
	if err != nil {
		return Time{}, err
	}

	return Time{
		hours,
		minutes,
		seconds,
	}, nil
}

func Errors() {
	time, err := ParseTime("14:07:33")

	if err != nil {
		panic(err)
	}

	fmt.Println(time.Hours)
	fmt.Println(time.Minutes)
	fmt.Println(time.Seconds)
}
