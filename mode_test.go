package mode_test

import (
	"testing"

	"github.com/lab47/mode"
)

func TestMode(t *testing.T) {
	if !mode.Debug() {
		t.Fatal("debug was not automatically enabled")
	}
}
