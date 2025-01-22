package rover_test

import (
	"github.com/rusticov/mars-rover-go-kata/internal/rover"
	"testing"
)

func TestMovingRover(t *testing.T) {

	testCases := []struct {
		name           string
		commands       string
		expectedOutput string
	}{
		{name: "should face north at origin when not moved", commands: "", expectedOutput: "0:0:N"},
		{name: "should face east at origin after rotating right once", commands: "R", expectedOutput: "0:0:E"},
		{name: "should face east at origin after rotating right twice", commands: "RR", expectedOutput: "0:0:S"},
		{name: "should face west at origin after rotating right 3 times", commands: "RRR", expectedOutput: "0:0:W"},
		{name: "should face back to north at origin after rotating right 4 times", commands: "RRRR", expectedOutput: "0:0:N"},
		{name: "should face west at origin after rotating left once", commands: "L", expectedOutput: "0:0:W"},
		{name: "should face south at origin after rotating left twice", commands: "LL", expectedOutput: "0:0:S"},
		{name: "should face east at origin after rotating left 3 times", commands: "LLL", expectedOutput: "0:0:E"},
		{name: "should face north at origin after rotating left 4 times", commands: "LLLL", expectedOutput: "0:0:N"},
		{name: "move north 1 time", commands: "M", expectedOutput: "0:1:N"},
		{name: "move north to top of the grid", commands: "MMMMMMMMM", expectedOutput: "0:9:N"},
		{name: "move north 10 times wraps location back to the start", commands: "MMMMMMMMMM", expectedOutput: "0:0:N"},
		{name: "move south moves down the grid", commands: "MMMLLM", expectedOutput: "0:2:S"},
		{name: "move south at bottom wraps location back to the top", commands: "LLM", expectedOutput: "0:9:S"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			movingRover := rover.New()

			output := movingRover.ExecuteCommands(testCase.commands)

			if output != testCase.expectedOutput {
				t.Errorf("expected %s but got %s", testCase.expectedOutput, output)
			}
		})
	}
}
