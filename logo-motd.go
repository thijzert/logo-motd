package main

import (
	"flag"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

var pixels = map[fourB]string{
	fourB{false, false, false, false}: " ",      // space
	fourB{true, false, false, false}:  "\u2598", // quadrant upper left
	fourB{false, true, false, false}:  "\u259d", // quadrant upper right
	fourB{true, true, false, false}:   "\u2580", // upper half block
	fourB{false, false, true, false}:  "\u2596", // quadrant lower left
	fourB{true, false, true, false}:   "\u258c", // left half block
	fourB{false, true, true, false}:   "\u259e", // ur/ll diagonal
	fourB{true, true, true, false}:    "\u259b", // inv. quadrant lower right
	fourB{false, false, false, true}:  "\u2597", // quadrant lower right
	fourB{true, false, false, true}:   "\u259a", // ul/lr diagonal
	fourB{false, true, false, true}:   "\u2590", // right half block
	fourB{true, true, false, true}:    "\u259c", // inv. quadrant lower left
	fourB{false, false, true, true}:   "\u2584", // lower half block
	fourB{true, false, true, true}:    "\u2599", // inv. quadrant upper right
	fourB{false, true, true, true}:    "\u259f", // inv quadrant upper left
	fourB{true, true, true, true}:     "\u2588"} // full block

var output_file = ""

func main() {
	flag.StringVar(&output_file, "o", "", "Output file name")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalf("Usage: %s FILE", os.Args[0])
	}

	var err error
	var out io.WriteCloser = os.Stdout
	logger := log.New(os.Stderr, "MOTD", log.Llongfile)
	exit := 0

	if output_file != "" {
		out, err = os.Create(output_file)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range args {
		if len(args) > 1 {
			out.Write([]byte(file + "\n"))
		}

		f, err := os.Open(file)
		if err != nil {
			logger.Print(err)
			exit = 1
			continue
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			logger.Print(err)
			exit = 1
			continue
		}

		size := img.Bounds()
		for y := size.Min.Y; y < size.Max.Y; y += 2 {
			for x := size.Min.X; x < size.Max.X; x += 2 {
				ff := fourB{
					isBlack(img.At(x, y)),
					isBlack(img.At(x+1, y)),
					isBlack(img.At(x, y+1)),
					isBlack(img.At(x+1, y+1))}

				out.Write([]byte(pixels[ff]))
			}
			out.Write([]byte("\n"))
		}
	}

	os.Exit(exit)
}

type fourB struct {
	A, B, C, D bool
}

func isBlack(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	if r > 0x1000 {
		return false
	}
	if g > 0x1000 {
		return false
	}
	if b > 0x1000 {
		return false
	}
	return true
}
