package main

import(
	"fmt"
	"os"
	"strings"
)

var asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers = "0123456789"
var mathSymbols = []string{"+", "-", "*", "**", "/", "(", ")"}


type returnState struct {
	nextState func(string) *returnState
	append    bool
	complete  bool
	increment bool
}


// GenericState is
func GenericState(char string, mapping map[string] *returnState) *returnState{

	// If the current state maps letters and char is a letter
	if _, ok := mapping["letters"]; ok {
		if strings.Contains(asciiChars, char) {
			return mapping["letters"]
		}
	}

	// If the current state maps numbers and char is a number
	if _, ok := mapping["numbers"]; ok {
		if strings.Contains(numbers, char) {
			return mapping["numbers"]
		}
	}

	// If the current state maps mathematical symbols
	// and char is a mathematical symbol
	if _, ok := mapping["math"]; ok {
		for _, sym := range mathSymbols {
			if sym == char {
				return mapping["numbers"]
			}
		}
	}

	// If char is not in the map, it is illegal
	if _, ok := mapping[char]; ok != true {
		fmt.Println("Illegal character:", char)
		os.Exit(1)
	}

	return mapping[char]
}

func StartState(char string) *returnState{

	mapping := map[string] *returnState{
		"numbers": {NumPreDotState, true, true, true},
		"+": {StartState, true, true, true},
	}

	return GenericState(
		char,
		mapping,
	)
}

func NumPreDotState(char string) *returnState{

	mapping := map[string] *returnState{
		"numbers": {NumPreDotState, true, false, true},
		"math": {StartState, true, true, true},
	}

	return GenericState(
		char,
		mapping,
	)
}

func NumPostDotState(char string) *returnState{

	mapping := map[string] *returnState{
		"numbers": {StartState, true, true, true},
		"+": {StartState, true, true, true},
	}

	return GenericState(
		char,
		mapping,
	)
}


func main() {

	var char string
	var inputSlice []string
	var stack []string
	var output []string

	input_string := "2+2"
	for _, char := range input_string {
		char := string(char)
		inputSlice = append(inputSlice, char)
	}

	state := StartState

	idx := 0
	for {
		char = inputSlice[idx]
		currentState := state(char)

		if currentState.append {
			stack = append(stack, char)
		}

		if currentState.increment {
			idx++
		}

		if currentState.complete {
			output = append(output, strings.Join(stack, ""))
			stack = nil
		}

		state = currentState.nextState

		if idx == len(inputSlice) {
			break
		}

	}

	fmt.Println(output)


}
