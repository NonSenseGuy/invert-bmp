package main

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/image/bmp"
)

type rgb struct {
	r uint8
	g uint8
	b uint8
}

const (
	N_MUESTRAS = 100
)

var (
	version       = 1
	inputImgPath  = filepath.FromSlash("./img/c.bmp")
	outputImgPath = filepath.FromSlash("./img/inverted_img.bmp")
)

func invert(t int, in, out string) error {
	dat, err := os.Open(in)
	if err != nil {
		return err
	}

	defer func() {
		errClose := dat.Close()
		if err == nil {
			err = errClose
		}
	}()

	img, err := bmp.Decode(dat)
	if err != nil {
		return err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	rgbArr0 := makeArray(height, width, img)

	rgbArr := makeArray(height, width, img)

	return writeImg(t, height, width, rgbArr0, rgbArr)
}

func makeArray(height, width int, img image.Image) [][]rgb {
	ImRGB0 := [][]rgb{}

	for r := 0; r < height; r++ {
		row := []rgb{}
		for c := 0; c < width; c++ {
			rx, gx, bx, ax := img.At(c, r).RGBA()
			r, g, b, _ := rx>>8, gx>>8, bx>>8, ax>>8

			temp := rgb{
				r: uint8(r),
				g: uint8(g),
				b: uint8(b),
			}
			row = append(row, temp)
		}
		ImRGB0 = append(ImRGB0, row)
	}

	return ImRGB0
}

func writeImg(version, height, width int, rgbArr0, rgbArr [][]rgb) error {

	for n := 0; n < N_MUESTRAS; n++ {
		start := time.Now()
		switch version {
		case 1:
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {

					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b
				}
			}
		default:
			break
		}

		stop := time.Now()
		_ = stop.Sub(start)

		// fmt.Println(elapsed)

	}

	//Write new img

	f, err := os.Create(outputImgPath)
	if err != nil {
		return err
	}

	upLeft := image.Point{0, 0}
	upRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, upRight})

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			color := color.RGBA{rgbArr[r][c].r, rgbArr[r][c].g, rgbArr[r][c].b, 255}
			img.Set(c, r, color)
		}
	}

	return bmp.Encode(f, img)
}

func main() {
	invert(1, inputImgPath, outputImgPath)
	// dat, _ := os.Open(inputImgPath)

	// img, _ := bmp.Decode(dat)
	// width := img.Bounds().Dx()
	// height := img.Bounds().Dy()
	// rgbArr := makeArray(height, width, img)

	// upLeft := image.Point{0, 0}
	// upRight := image.Point{width, height}
	// nImg := image.NewRGBA(image.Rectangle{upLeft, upRight})

	// for r := 0; r < height; r++ {
	// 	for c := 0; c < width; c++ {
	// 		color := color.RGBA{rgbArr[r][c].r, rgbArr[r][c].g, rgbArr[r][c].b, 255}
	// 		nImg.Set(c, r, color)
	// 	}
	// }

	// f, _ := os.Create(outputImgPath)

	// bmp.Encode(f, nImg)

}
