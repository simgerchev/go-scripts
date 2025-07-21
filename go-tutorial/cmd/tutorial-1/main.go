package main

import (
	"errors"
	"fmt"
)

func main() {
	numerator := 100
	denumerator := 10
	result, remainder, err := intDivision(numerator, denumerator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("%v", result)
	} else {
		fmt.Printf("%v.%v", result, remainder)
	}

}

func intDivision(numerator int, denumerator int) (int, int, error) {
	var err error
	if denumerator == 0 {
		err = errors.New("invalid input: a must be non-negative and b must be positive")
		return 0, 0, err
	}
	result := numerator / denumerator
	remainder := numerator % denumerator
	return result, remainder, err
}
