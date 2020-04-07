// +build windows plan9 js

package uumask

// Umask is not supported on these platforms
func Umask() uint32 {
	return 0000
}
