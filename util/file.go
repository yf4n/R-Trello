package util

import (
	"os"
)

func WriteFile(path string, content string) {
	f, err := os.Create(path)
	CheckError(err)
	defer f.Close()
	f.WriteString(content)
	CheckError(err)
	f.Sync()
}
