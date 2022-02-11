package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	message := GetMessage()
	if message != "Hello 🗺️!" {
		t.Error()
	}
}
