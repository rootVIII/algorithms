package main

// rootVIII

// SubImage example: Creates 18, 128x128-pixel
// sub-images/PNGs taken from atlas.png.

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

func writePNG(img draw.Image, path string) {
	buf := &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		err = ioutil.WriteFile(path, buf.Bytes(), 0600)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func main() {
	imgRaw, _ := os.Open("atlas.png")
	defer imgRaw.Close()
	atlas, _, _ := image.Decode(imgRaw)

	img := image.NewRGBA(image.Rect(0, 0, 768, 384))
	draw.Draw(img, img.Bounds(), atlas, image.Point{}, draw.Src)

	rows, cols := 3, 6
	for r := 0; r < rows; r++ {
		for c, pixels := 0, 0; c < cols; c, pixels = c+1, pixels+128 {
			subImg := img.SubImage(image.Rect(pixels, r*128, pixels+128, (r+1)*128)).(*image.RGBA)
			writePNG(subImg, fmt.Sprintf("%d-%d.png", r+1, c+1))
		}
	}

	/*
		- row 0
		0, 0, 128, 128
		128, 0, 256, 128
		256, 0, 384, 128
		384, 0, 512, 128
		512, 0, 640, 128
		640, 0, 768, 128

		- row 1
		0, 128, 128, 256
		128, 128, 256, 256
		256, 128, 384, 256
		384, 128, 512, 256
		512, 128, 640, 256
		640, 128, 768, 256


		- row 2
		0, 256, 128, 384
		128, 256, 256, 384
		256, 256, 384, 384
		384, 256, 512, 384
		512, 256, 640, 384
		640, 256, 768, 384
	*/
}
