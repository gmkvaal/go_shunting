package main

import (
	"fmt"
	//"strings"
	"github.com/gmkvaal/go-shunting/states"
	"strings"
)

// Formats inputString into individual tokens
func main() {

	var char string
	var inputSlice []string
	var stack []string
	var output []string

	inputString := "2.2+3.2**2*1//2"
	for _, char := range inputString {
		char := string(char)
		inputSlice = append(inputSlice, char)
	}

	state := states.StartState

	idx := 0
	for {
		char = inputSlice[idx]
		currentState, err:= state(char)
		if err != nil {
			fmt.Print(err)
		}

		if currentState.Append {
			stack = append(stack, char)
		}

		if currentState.Increment {
			idx++
		}

		if currentState.Complete {
			output = append(output, strings.Join(stack, ""))
			stack = nil
		}

		if idx == len(inputSlice) {
			if !currentState.Complete {
				output = append(output, strings.Join(stack, ""))
			}
			break
		}

		state = currentState.NextState

	}

	fmt.Println(output)
}
