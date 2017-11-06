package main

import(
	"fmt"
	"os"
	"strings"
)

var asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers = "0123456789"



type returnState struct {
	nextState func(string) returnState
	append    bool
	complete  bool
	increment bool
}

func SuperState(char string, mapping map[string] returnState) returnState{

	// If the current state maps letters and the char is a letter
	if _, ok := mapping["letters"]; ok {
		if strings.Contains(asciiChars, char) {
			return mapping["letters"]
		}
	}

	// If the current state maps numbers and the char is a number
	if _, ok := mapping["numbers"]; ok {
		if strings.Contains(numbers, char) {
			return mapping["numbers"]
		}
	}

	// If char is not in the map, it is illegal
	if _, ok := mapping[char]; ok != true {
		fmt.Println("Illegal character:", char)
		os.Exit(1)
	}

	return mapping[char]
}

func StartState(char string) returnState{

	mapping := map[string] returnState{
		"letters": {StartState, true, false, false},
		"+": {StartState, true, false, false},
	}

	return SuperState(
		char,
		mapping,
	)
}

func main() {


	char := "b"


	state := StartState(char)


	fmt.Println(state.append)
}
