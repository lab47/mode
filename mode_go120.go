//go:build !go1.21
// +build !go1.21

package mode

import "flag"

func inTesting() bool {
	return flag.Lookup("test.v") != nil
}
