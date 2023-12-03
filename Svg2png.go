package svg2png

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func Svg2png(options Options) {
	if !normalizeOptions(&options) {
		fmt.Println("Invalid options")
		return
	}

	in, err := os.Open(options.InputFilePath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	icon, err := oksvg.ReadIconStream(in)
	if err != nil {
		panic(err)
	}

	// Calculate the aspect ratio
	svgWidth := icon.ViewBox.W
	svgHeight := icon.ViewBox.H
	aspectRatio := svgWidth / svgHeight

	var w, h float64
	if aspectRatio > 1 { // width is greater than height
		w = float64(options.MaxWidth)
		h = float64(options.MaxWidth) / aspectRatio
	} else {
		w = float64(options.MaxHeight) * aspectRatio
		h = float64(options.MaxHeight)
	}

	// Set the target dimensions while preserving aspect ratio
	icon.SetTarget(0, 0, w, h)
	rgba := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	icon.Draw(rasterx.NewDasher(int(w), int(h), rasterx.NewScannerGV(int(w), int(h), rgba, rgba.Bounds())), 1)

	out, err := os.Create(options.OutputFilePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		panic(err)
	}
}
