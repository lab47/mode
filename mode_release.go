//go:build release
// +build release

package mode

func init() {
	DefaultMode = ReleaseMode
}
