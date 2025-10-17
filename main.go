package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {

	args := &Args{}

	app := &cli.App{
		Name:  "csv2md",
		Usage: "convert csv to markdown",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "padding",
				Aliases:     []string{"p"},
				Value:       2,
				Usage:       "The number of spaces to add between table cells and column dividers.",
				Destination: &args.Pad,
			},
			&cli.StringFlag{
				Name:        "delimiter",
				Aliases:     []string{"d"},
				Value:       ",",
				Usage:       "CSV delimiter, expected values: ',', ';'.",
				Destination: &args.Delim,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "",
				Usage:       "Write output to file",
				Destination: &args.OutputFile,
			},
		},
		Action: func(c *cli.Context) error {
			args.Files = c.Args().Slice()
			if err := args.ValidateAll(); err != nil {
				return err
			}

			if err := ConvertAll(args); err != nil {
				return err
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
