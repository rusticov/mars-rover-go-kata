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
