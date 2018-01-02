package testutils

import "runtime"

var executablePath string

func init() {
	if runtime.GOOS == "darwin" {
		executablePath = "build/dist/darwin/ironman"
	} else {
		executablePath = "build/dist/linux/ironman"
	}
}

//ExecutablePath returns the build executable path based on the os
func ExecutablePath() string {
	return executablePath
}
