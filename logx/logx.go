package lgox

import "os"

var logFilePtr = os.Stderr

func SetLogFilePtr(fx *os.File) {
	logFilePtr = fx
}

func GetLogFilePtr() (fx *os.File) {
	return logFilePtr
}
