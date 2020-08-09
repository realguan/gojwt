package utils

import (
	"os"
)

// FileExist ..
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}