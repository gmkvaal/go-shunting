package main

import(
	"fmt"
	"os"
	"reflect"
)

// var mapping map[string]

type returnState struct {
	nextState func(string) returnState
	append    bool
	complete  bool
	increment bool
}




func SuperState(char string, mapping map[string] returnState) returnState{

	// If char is not in the map, it is illegal
	if _, ok := mapping[char]; ok != true {
		fmt.Println("Illegal character:", char)
		os.Exit(1)
	}

	fmt.Println(reflect.TypeOf(mapping[char]))

	return(mapping[char])
}

func TestState(char string) returnState{

	mapping := map[string] returnState{
		"b": {TestState, false, false, false},
	}

	return SuperState(
		char,
		mapping,
	)
}

func main() {


	char := "b"


	state := TestState(char)


	fmt.Println(state)
}
