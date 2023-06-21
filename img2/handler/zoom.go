package handler

import (
	"errors"
	"fmt"
	"image"

	"github.com/urfave/cli/v2"
	"golang.org/x/image/draw"
)

func CheckResize(c *cli.Context) bool {
	w := c.Int("width")
	h := c.Int("height")
	return w+h > 0
}

func Resize(c *cli.Context, img image.Image) (image.Image, error) {
	newWidth := c.Int("width")
	newHeight := c.Int("height")

	if newWidth+newHeight == 0 {
		return nil, errors.New("0 value of newWidth and newHeight")
	}
	size := img.Bounds().Size()
	if newWidth == 0 || newHeight == 0 {
		if newHeight == 0 {
			newHeight = newWidth * size.Y / size.X
		}
	}
	fmt.Println(size, ">", image.Point{X: newWidth, Y: newHeight})
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return newImg, nil
}

func Cut(c *cli.Context, img image.Image) (image.Image, error) {
	return nil, nil
}
