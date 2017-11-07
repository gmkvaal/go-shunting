package states

import (
	"strings"
	"fmt"
	"os"
)

var asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var mathSymbols = []string{"+", "-", "*", "**", "/", "(", ")"}

// Struct field containing the next state,
// if Append char to stack, if the state is
// complete, or if Increment (read next char).
type ReturnState struct {
	NextState func(string) *ReturnState
	Append    bool
	Complete  bool
	Increment bool
}

// GenericState interprets the mappings for the current state in relation to
// the current char and returns the correct mapping
func GenericState(char string, mapping map[string]*ReturnState, illegals []string) *ReturnState {

	// If the current state maps letters and char is a letter.
	if _, ok := mapping["letters"]; ok {
		if strings.Contains(asciiChars, char) {
			return mapping["letters"]
		}
	}

	// If the current state maps digits and char is a number.
	if _, ok := mapping["digits"]; ok {
		if strings.Contains(digits, char) {
			return mapping["digits"]
		}
	}

	// If the current state maps mathematical symbols
	// and char is a mathematical symbol.
	if _, ok := mapping["math"]; ok {
		for _, sym := range mathSymbols {
			if sym == char {
				return mapping["math"]
			}
		}
	}

	// If char is in illegals, it is illegal.
	for _, illegal := range illegals {
		if illegal == char {
			fmt.Println("Character in illegals:", char)
			os.Exit(1)
		}
	}

	// If char is not in the map, check if a 'default' mapping exists.
	// If not, char is illegal
	if _, ok := mapping[char]; ok != true {
		if _, ok := mapping["default"]; ok {
			return mapping["default"]
		} else {
			fmt.Println("Illegal character:", char)
			os.Exit(1)
		}
	}

	return mapping[char]
}

// Start state is the initial state. Maps to next state
// without any restrictions on next char.
func StartState(char string) *ReturnState {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"digits": {NumPreDotState, true, false, true},
		"math":   {SymState, false, false, false},
		".":      {NumPostDotState, true, false, true},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}

// SymState governs the tokenizing of mathematical symbols.
// Returns call to GenericState with char, mapping and illegals.
func SymState(char string) *ReturnState {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"(": {StartState, true, true, true},
		")": {StartState, true, true, true},
		"%": {StartState, true, true, true},
		"-": {StartState, true, true, true},
		"+": {StartState, true, true, true},
		"*": {MulState, true, false, true},
		"/": {DivState, true, false, true},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}

// MulState checks if '*' is preceded by another '*' to form '**'
// Returns call to GenericState with char, mapping and illegals
func MulState(char string) *ReturnState {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"*":       {StartState, true, true, true},
		"default": {StartState, false, true, false},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}

// DivState checks if '/' is preceded by another '/' to form '//'
// Returns call to GenericState with char, mapping and illegals
func DivState(char string) *ReturnState {

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"/":       {StartState, true, true, true},
		"default": {StartState, false, true, false},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}

// NumPreDotState governs the tokenization of digits prior to decimals.
// Returns call to GenericState with char, mapping and illegals
func NumPreDotState(char string) *ReturnState {

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"digits": {NumPreDotState, true, false, true},
		"math":   {SymState, false, true, false},
		".":      {NumPostDotState, true, false, true},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}

// NumPreDotState governs tokenization of decimal digits.
// Returns call to GenericState
func NumPostDotState(char string) *ReturnState {

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"digits": {StartState, true, true, true},
		"math":   {SymState, false, true, false},
		"+":      {StartState, false, true, true},
	}

	return GenericState(
		char,
		mapping,
		illegals,
	)
}
