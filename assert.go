package mode

// If val is not true and the mode is debug, panic with the given message.
func Assert(val bool, reason string) {
	if Debug() && !val {
		panic(reason)
	}
}
