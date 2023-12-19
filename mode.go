package mode

import (
	"sync"
)

const (
	ReleaseMode = "release"
	DebugMode   = "debug"
	TestMode    = "test"
)

var (
	// DefaultMode is the mode that the program should become if there
	// is no mode yet set.
	DefaultMode string

	isDebug     bool
	currentMode string
)

var resolve sync.Once

// Mode returns the current mode.
func Mode() string {
	resolve.Do(func() {
		if DefaultMode != "" {
			currentMode = DefaultMode
		} else if inTesting() {
			// We don't want to use testing.Testing() here because that will
			// import the testing package into all users of mode which we
			// don't want.
			currentMode = TestMode
		} else {
			// Match the way most folks use go, ie, the default is release
			// mode.
			currentMode = ReleaseMode
		}

		if currentMode != ReleaseMode {
			isDebug = true
		}
	})

	return currentMode
}

// Debug returns true if the current mode is not release.
func Debug() bool {
	// to force the current mode to resolve
	_ = Mode()

	return isDebug
}
