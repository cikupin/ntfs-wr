package pkg

import "runtime"

func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}
