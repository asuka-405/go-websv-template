package libfs

import (
	"fmt"
	"os"
)

func UnsafeFileReader(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Err reading file: %v", err)
		return ""
	}
	return string(data)
}
