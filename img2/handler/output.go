package handler

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func Encode(f *os.File, format string, img image.Image) error {
	switch format {
	case "jpeg":
		return jpeg.Encode(f, img, nil)
	case "png":
		return png.Encode(f, img)
	default:
		return errors.New("unknown format " + format)
	}
}
