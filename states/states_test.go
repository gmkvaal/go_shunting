package states

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

	mapping := map[string]*ReturnState{
		"digits": {numPreDotState, true, false, true},
		"math":   {symState, false, false, false},
		".":      {numPostDotState, true, false, true},
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
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[mapKey].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[mapKey].Append, &output.Append)
		assert.Equal(t, &mapping[mapKey].Complete, &output.Complete)
		assert.Equal(t, &mapping[mapKey].Increment, &output.Increment)
	}
}

func TestsymState(t *testing.T) {

	mapping := map[string]*ReturnState{
		"(": {StartState, true, true, true},
		")": {StartState, true, true, true},
		"%": {StartState, true, true, true},
		"-": {StartState, true, true, true},
		"+": {StartState, true, true, true},
		"*": {mulState, true, false, true},
		"/": {divState, true, false, true},
	}

	for char := range mapping {
		output := symState(char)
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[char].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[char].Append, &output.Append)
		assert.Equal(t, &mapping[char].Complete, &output.Complete)
		assert.Equal(t, &mapping[char].Increment, &output.Increment)
	}
}

func TestMulState(t *testing.T) {

	mapping := map[string]*ReturnState{
		"+": {StartState, false, true, false},
		"*": {StartState, true, true, true},
		"1": {StartState, false, true, false},
	}

	for char := range mapping {
		output := mulState(char)
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[char].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[char].Append, &output.Append)
		assert.Equal(t, &mapping[char].Complete, &output.Complete)
		assert.Equal(t, &mapping[char].Increment, &output.Increment)
	}
}

func TestDivState(t *testing.T) {

	mapping := map[string]*ReturnState{
		"+": {StartState, false, true, false},
		"/": {StartState, true, true, true},
		"1": {StartState, false, true, false},
		".": {StartState, false, true, false},

	}

	for char := range mapping {
		output := divState(char)
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[char].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[char].Append, &output.Append)
		assert.Equal(t, &mapping[char].Complete, &output.Complete)
		assert.Equal(t, &mapping[char].Increment, &output.Increment)
	}
}

func TestNumPreDotState(t *testing.T) {

	mapping := map[string]*ReturnState{
		"+": {symState, false, true, false},
		".": {numPostDotState, true, false, true},
		"1": {numPreDotState, true, false, true},
	}

	for char := range mapping {
		output := numPreDotState(char)
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[char].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[char].Append, &output.Append)
		assert.Equal(t, &mapping[char].Complete, &output.Complete)
		assert.Equal(t, &mapping[char].Increment, &output.Increment)
	}
}

func TestNumPostDotState(t *testing.T) {

	mapping := map[string]*ReturnState{
		"+": {symState, false, true, false},
		"1": {numPostDotState, true, false, true},
	}

	for char := range mapping {
		output := numPostDotState(char)
		// Asserting that correct NextState is returned by comparing the function name
		// as assertion of type func is not allowed
		assert.Equal(t, GetFuncName(mapping[char].NextState), GetFuncName(output.NextState))
		assert.Equal(t, &mapping[char].Append, &output.Append)
		assert.Equal(t, &mapping[char].Complete, &output.Complete)
		assert.Equal(t, &mapping[char].Increment, &output.Increment)
	}
}

