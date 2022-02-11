package solution

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	message := GetMessage()
	if message != "Hello 🗺️ !" {
		t.Error()
	}
}

func TestGetMessage(t *testing.T) {
	expected := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	msg := GetMessage()

	if !strings.EqualFold(msg, expected) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}