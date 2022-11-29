package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	var outFormat = flag.String("t", "jpeg", "specify output format")
	var outName = flag.String("o", "test.jpg", "specify output file name")
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input format = %s\n", kind)
	fmt.Printf("Output format = %s\n", *outFormat)

	f, err := os.Create(*outName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file: %s\n", *outName)
	}

	if err := convert(img, *outFormat, f); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
}

func convert(img image.Image, outFormat string, out io.Writer) error {
	switch outFormat {
	case "jpeg":
		return toJPEG(img, out)
	case "png":
		return toPNG(img, out)
	case "gif":
		return toGIF(img, out)
	default:
		return errors.New("invalid output format")
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
}
