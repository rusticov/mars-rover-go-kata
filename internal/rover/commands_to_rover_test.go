package rover_test

import (
	"github.com/rusticov/mars-rover-go-kata/internal/rover"
	"testing"
)

func TestMovingRover(t *testing.T) {

	t.Run("should face north at origin when not moved", func(t *testing.T) {
		movingRover := rover.New()

		location := movingRover.ExecuteCommands("")

		if location != "0:0:N" {
			t.Errorf("expected 0:0:N but got %s", location)
		}
	})

	t.Run("should face east at origin when rotating right", func(t *testing.T) {
		movingRover := rover.New()

		location := movingRover.ExecuteCommands("R")

		if location != "0:0:E" {
			t.Errorf("expected 0:0:E but got %s", location)
		}
	})
}
