package libcache

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadCredentials() (string, string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	sp_util_creds := filepath.Join(homeDir, ".sp-util-creds")
	if _, err := os.Stat(sp_util_creds); os.IsNotExist(err) {
		panic("Credentials file not found")
	}
	content, err := os.ReadFile(sp_util_creds)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) != 2 {
		panic("Invalid credentials file")
	}
	return lines[0], lines[1]
}
