package utils

import "path/filepath"

func SanitizePath(path string) string {
	return filepath.Clean(path)
}
