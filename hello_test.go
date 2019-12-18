package m1

import (
	"testing"
)

func TestM1(t *testing.T) {
	expected := "hello"
	got := Hello()
	if got != expected {
		t.Errorf("expected: %v\tgot: %v\n", expected, got)
	}
}
