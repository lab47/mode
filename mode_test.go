//go:build test
// +build test

package mode

import "testing"

func init() {
	if testing.Testing() {
		DefaultMode = TestMode
	}
}
