package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"runtime"
	"reflect"
)


// GetFuncName returns the name of a function
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}


func TestStartState(t *testing.T) {

	mapping := map[string]*returnState{
		"digits": {NumPreDotState, true, false, true},
		"math":   {StartState, true, true, true},
		".":      {NumPostDotState, true, false, true},
	}

	inputOutputMap := map[string] string {
		"1": "digits",
		"+": "math",
		"-": "math",
		"*": "math",
		".": ".",
	}

	for char, mapKey := range inputOutputMap {
		output := StartState(char)
		// Asserting that correct nextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[mapKey].nextState), GetFuncName(output.nextState))
		assert.Equal(t, &mapping[mapKey].append, &output.append)
		assert.Equal(t, &mapping[mapKey].complete, &output.complete)
		assert.Equal(t, &mapping[mapKey].increment, &output.increment)
	}
}