package mode

func Assert(val bool, reason string) {
	if Debug() && !val {
		panic(reason)
	}
}
