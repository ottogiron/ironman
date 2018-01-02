package testutils

import "runtime"

var executablePath string

func init() {
	if runtime.GOOS == "darwin" {
		executablePath = "build/dist/darwin/ironman"
	} else if runtime.GOOS == "linux" {
		executablePath = "build/dist/linux/ironman"
	} else if runtime.GOOS == "windows" {
		executablePath = "build/dist/windows/ironman.exe"
	}
}

//ExecutablePath returns the build executable path based on the os
func ExecutablePath() string {
	return executablePath
}
