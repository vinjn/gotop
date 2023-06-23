//go:build !arm64
// +build !arm64

package logging

import (
	"os"
)

func StderrToLogfile(logfile *os.File) {
	// syscall.Dup2(int(logfile.Fd()), 2)
}
