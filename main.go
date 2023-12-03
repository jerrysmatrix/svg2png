package svg2png

func main() {
	options := Options{InputFilePath: "test.svg", MaxWidth: 512, MaxHeight: 512}
	Svg2png(options)
}
