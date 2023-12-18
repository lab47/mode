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
	DefaultMode string
	CurrentMode string
	IsDebug     bool
)

var resolve sync.Once

func Mode() string {
	resolve.Do(func() {
		if DefaultMode != "" {
			CurrentMode = DefaultMode
		} else if inTesting() {
			// We don't want to use testing.Testing() here because that will
			// import the testing package into all users of mode which we
			// don't want.
			CurrentMode = TestMode
		} else {
			// Match the way most folks use go, ie, the default is release
			// mode.
			CurrentMode = ReleaseMode
		}

		if CurrentMode != ReleaseMode {
			IsDebug = true
		}
	})

	return CurrentMode
}

func Debug() bool {
	// to force the current mode to resolve
	_ = Mode()

	return IsDebug
}
