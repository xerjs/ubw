package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/image/draw"
)

func Resize(c *cli.Context, img image.Image) (image.Image, error) {
	newWidth := c.Int("width")
	newHeight := c.Int("height")

	if newWidth+newHeight == 0 {
		return nil, errors.New("0 value of newWidth newHeight")
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

func main() {
	app := cli.NewApp()
	app.Name = "img2"
	app.Usage = "image tool"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Required: true,
			Usage:    "image file path",
		},
		&cli.IntFlag{
			Name:    "width",
			Aliases: []string{"w"},
		},
		&cli.IntFlag{
			Name: "height",
		},
		&cli.StringFlag{
			Name:     "output",
			Required: true,
		},
	}
	app.Action = func(c *cli.Context) error {
		f, err := os.Open(c.String("input"))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		img, formatName, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(formatName)
		img2, err := Resize(c, img)
		if err != nil {
			panic(err)
		}

		outFile, err := os.Create(c.String("output"))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		err = Encode(outFile, formatName, &img2)
		if err != nil {
			panic(err)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
