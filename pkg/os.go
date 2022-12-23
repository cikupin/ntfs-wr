package pkg

import "runtime"

// IsMacOS detects whether the OS is macOS or not
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}
