package svg2png

import "os"

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return err == nil // File exists and no other error occurred.
}
