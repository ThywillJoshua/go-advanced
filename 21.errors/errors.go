package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func parseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of components", input}
	}

	hour, err := strconv.Atoi(components[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
	}

	minute, err := strconv.Atoi(components[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute: %v", err), input}
	}

	second, err := strconv.Atoi(components[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", err), input}
	}

	if hour > 23 || hour < 0 {
		return Time{}, &TimeParseError{fmt.Sprintf("Hour out of range: 0 <= hour <= 23: %v", hour), input}
	}

	if minute > 59 || minute < 0 {
		return Time{}, &TimeParseError{fmt.Sprintf("Minute out of range: 0 <= minute <= 59: %v", minute), input}
	}

	if second > 59 || second < 0 {
		return Time{}, &TimeParseError{fmt.Sprintf("Minute out of range: 0 <= second <= 59: %v", second), input}
	}

	return Time{hour, minute, second}, nil
}

func main() {
}
