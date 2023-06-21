package handler

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func Encode(f *os.File, format string, img *image.Image) error {
	switch format {
	case "jpeg":
		return jpeg.Encode(f, *img, &jpeg.Options{Quality: 50})
	case "png":
		return png.Encode(f, *img)
	default:
		return errors.New("unknown format " + format)
	}
}
