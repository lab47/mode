## mode

Mode is a simple Golang package designed to allow code to run in debug, release, or test mode.

The idea is that code can conditional code using `mode.Debug()` and only perform
certain activities when the code is not in release mode.

### Specifying A Mode

To specify the mode to compile the code under, you use build tags.

To compile in debug mode, add `-tags debug` to `go build`.

To compile in release mode, add `-tags release` to `go build`.

#### Default Mode

By default, `mode` assumes the mode is `release` unless specified or detected
otherwise. This matches the default semantics of Go programs today.

### Test detection

Mode can automatically determine if the code is running via `go test`. 

Starting with Go 1.21, the `testing` package can safely be imported by non-test packages,
and so that is used to call `testing.Testing()`. Otherwise, at runtime, mode checks
for the test flags.
