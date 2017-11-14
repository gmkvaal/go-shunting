package states

import (
	"strings"
	"errors"
)
const (
	asciiChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "0123456789"
)

var mathSymbols = [7]string{"+", "-", "*", "**", "/", "(", ")"}

// Struct field containing the next state,
// if Append char to stack, if the state is
// complete, or if Increment (read next char).
type ReturnState struct {
	NextState func(string) (*ReturnState, error)
	Append    bool
	Complete  bool
	Increment bool
}

// genericState interprets the mappings for the current state in relation to
// the current char and returns the correct mapping.
func genericState(char string, mapping map[string]*ReturnState, illegals []string) (*ReturnState, error) {

	// If the current state maps letters and char is a letter.
	if _, ok := mapping["letters"]; ok {
		if strings.Contains(asciiChars, char) {
			return mapping["letters"], nil
		}
	}

	// If the current state maps digits and char is a number.
	if _, ok := mapping["digits"]; ok {
		if strings.Contains(digits, char) {
			return mapping["digits"], nil
		}
	}

	// If the current state maps mathematical symbols
	// and char is a mathematical symbol.
	if _, ok := mapping["math"]; ok {
		for _, sym := range mathSymbols {
			if sym == char {
				return mapping["math"], nil
			}
		}
	}

	// If char is in illegals, it is illegal.
	for _, illegal := range illegals {
		if illegal == char {
			return nil, errors.New("illegal char")
		}
	}

	// If char is not in the map, check if a 'default' mapping exists.
	// If not, char is illegal
	if _, ok := mapping[char]; ok != true {
		if _, ok := mapping["default"]; ok {
			return mapping["default"], nil
		} else {
			return nil, errors.New("illegal char")
		}
	}

	return mapping[char], nil
}

// Start state is the initial state. Maps to next state
// without any restrictions on next char.
func StartState(char string) (*ReturnState, error) {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"digits": {numPreDotState, true, false, true},
		"math":   {symState, false, false, false},
		".":      {numPostDotState, true, false, true},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}

// symState governs the tokenizing of mathematical symbols.
// Returns call to genericState with char, mapping and illegals.
func symState(char string) (*ReturnState, error) {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"(": {StartState, true, true, true},
		")": {StartState, true, true, true},
		"%": {StartState, true, true, true},
		"-": {StartState, true, true, true},
		"+": {StartState, true, true, true},
		"*": {mulState, true, false, true},
		"/": {divState, true, false, true},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}

// mulState checks if '*' is preceded by another '*' to form '**'
// Returns call to genericState with char, mapping and illegals
func mulState(char string) (*ReturnState, error) {

	illegals := []string{","}
	mapping := map[string]*ReturnState{
		"*":       {StartState, true, true, true},
		"default": {StartState, false, true, false},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}

// divState checks if '/' is preceded by another '/' to form '//'
// Returns call to genericState with char, mapping and illegals
func divState(char string) (*ReturnState, error) {

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"/":       {StartState, true, true, true},
		"default": {StartState, false, true, false},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}

// numPreDotState governs the tokenization of digits prior to decimals.
// Returns call to genericState with char, mapping and illegals
func numPreDotState(char string) (*ReturnState, error) {

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"digits": {numPreDotState, true, false, true},
		"math":   {symState, false, true, false},
		".":      {numPostDotState, true, false, true},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}

// numPreDotState governs tokenization of decimal digits.
// Returns call to genericState
func numPostDotState(char string) (*ReturnState, error){

	illegals := []string{}
	mapping := map[string]*ReturnState{
		"digits": {numPostDotState, true, false, true},
		"math":   {symState, false, true, false},
	}

	return genericState(
		char,
		mapping,
		illegals,
	)
}
