package libcache

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Initialize() {
	// Initialize the cache
	cacheDir := filepath.Join(os.TempDir(), "cache")
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cache initialized at: ", cacheDir)
}
