package svg2png

import (
	"fmt"
	"path/filepath"
)

func normalizeOptions(options *Options) bool {
	if options.InputFilePath == "" || !fileExists(options.InputFilePath) {
		return false
	}

	if options.MaxHeight < 1 {
		options.MaxHeight = 512
	}

	if options.MaxWidth < 1 {
		options.MaxWidth = 512
	}

	if options.OutputFileName == "" {
		options.OutputFileName = getFileNameWithoutExtension(options.InputFilePath)
	}

	if options.OutputDirectory == "" {
		options.OutputDirectory = filepath.Dir(options.InputFilePath)
	}

	options.OutputFileName = options.OutputFileName + ".png"
	options.OutputFilePath = filepath.Join(options.OutputDirectory, options.OutputFileName)

	fmt.Println("options.InputFilePath", options.InputFilePath)
	fmt.Println("options.MaxHeight", options.MaxHeight)
	fmt.Println("options.MaxWidth", options.MaxWidth)
	fmt.Println("options.OutputDirectory", options.OutputDirectory)
	fmt.Println("options.OutputFileName", options.OutputFileName)
	fmt.Println("options.OutputFilePath", options.OutputFilePath)

	return true
}
