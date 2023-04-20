package main

import (
	"fmt"
	"os"

	"github.com/muhammad-asn/igor/utils"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "identicon-generator"
	app.Usage = "generate identicon images"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "num-images, n",
			Value: 1,
			Usage: "number of images to generate",
		},
	}
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	if c.NumFlags() < 0 {
		cli.ShowAppHelp(c)
		os.Exit(0)
	}

	numImages := c.Int("num-images")

	for i := 0; i < numImages; i++ {
		filename := fmt.Sprintf("random_%d.png", i+1)
		err := utils.GenerateImage(filename)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Generated image %d\n", numImages)

	return nil
}
