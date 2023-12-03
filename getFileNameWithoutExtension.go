package svg2png

import (
	"path/filepath"
	"strings"
)

func getFileNameWithoutExtension(filePath string) string {
	filename := filepath.Base(filePath)
	return strings.TrimSuffix(filename, ".svg")
}
