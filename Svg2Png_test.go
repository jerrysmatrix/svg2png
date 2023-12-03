package svg2png

import (
	"os"
	"testing"
)

func TestSvg2png(t *testing.T) {
	options := Options{
		InputFilePath: "test.svg",
		MaxWidth:      512,
		MaxHeight:     512,
	}
	Svg2png(&options)
	outputFilePath := "test.png"

	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Errorf("Svg2png did not create the expected output file: %s", outputFilePath)
	}

	defer os.Remove(outputFilePath)
}
