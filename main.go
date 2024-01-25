package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	filename := "test.txt"
	app := &cli.App{
		Name:  "gwc",
		Usage: "Go implementation of the Unix wc command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lines",
				Aliases:     []string{"l"},
				Usage:       "displays number of lines in the file",
				Destination: &filename,
				Action: func(cCtx *cli.Context, s string) error {
					countString, err := lineCount(cCtx, s)
					fmt.Println((countString))
					return err
				},
			},
			&cli.StringFlag{
				Name:        "words",
				Aliases:     []string{"w"},
				Usage:       "displays the number of words in the file",
				Destination: &filename,
				Action: func(cCtx *cli.Context, s string) error {
					countString, err := wordCount(cCtx, s)
					fmt.Println((countString))
					return err
				},
			},
			&cli.StringFlag{
				Name:    "bytes",
				Aliases: []string{"c"},
				Usage:   "displays the count of bytes in the file", Destination: &filename,
				Action: func(cCtx *cli.Context, s string) error {
					countString, err := byteCount(cCtx, s)
					fmt.Println(countString)
					return err
				},
			},
			&cli.StringFlag{
				Name:        "char",
				Aliases:     []string{"m"},
				Usage:       "displays count of characters from a file",
				Destination: &filename,
				Action: func(cCtx *cli.Context, s string) error {
					countString, err := byteCount(cCtx, s)
					fmt.Println(countString)
					return err
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NumFlags() == 0 {
				fmt.Println("Support for default action to be added soon. Please use flags to indicate intended action. Run 'gwc help' for a list of available flags")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
