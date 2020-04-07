// +build !windows,!plan9,!js

package uumask

import "golang.org/x/sys/unix"

// Determine umask
// To determine umask, one must change it.
// Leaving umask changed even for a short time could have consequences,
// so it should be called at a time when there can be no file or directory creations.
func Umask() uint32 {
	// Change umask to get current value
	var original = unix.Umask(0000)

	// Restore umask
	unix.Umask(original)

	// Return original value
	return uint32(original)
}
