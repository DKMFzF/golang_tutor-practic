package abs

import "testing"

func TestAbs(t *testing.T) {
	if got := Abs(-10); got != 10 {
		t.Errorf("Abs(-10) -> %v", got)
	}
}
