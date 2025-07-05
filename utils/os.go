package utils

import (
	"fmt"
	"os"
	"runtime"
)

func CheckOSSupported() {
	if runtime.GOOS == "windows" {
		fmt.Fprintln(os.Stderr, "❌ Windows is not supported currently.")
		os.Exit(1)
	}
}
