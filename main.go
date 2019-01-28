package main

import (
	"errors"
	"insertnewline/actions/insertnl"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "insertnewline"
	app.Usage = "insert new line free position your text"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "source test file.",
		},
		cli.IntFlag{
			Name:  "num, n",
			Usage: "insert new line position num(1>n)",
		},
		cli.BoolFlag{
			Name:  "down, d",
			Usage: "use string down",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("file") == "" || c.Int("num") == 0 {
			return errors.New("invalid args!")
		}
		return insertnl.Action(c.String("file"), c.Int("num"), c.Bool("down"))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
