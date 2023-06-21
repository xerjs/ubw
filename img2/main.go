package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"ubw/img2/handler"

	"github.com/urfave/cli/v2"
)

func AppAction(c *cli.Context) error {
	f, err := os.Open(c.String("input"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, imgFormat, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(imgFormat)
	img2, err := handler.Resize(c, img)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(c.String("output"))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = handler.Encode(outFile, imgFormat, &img2)
	if err != nil {
		panic(err)
	}
	return nil
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
	app.Action = AppAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
