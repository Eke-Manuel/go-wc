package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "gwc",
		Usage:                  "Go implementation of the Unix wc command",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "lines",
				Aliases: []string{"l"},
				Usage:   "displays number of lines in the file",
				Action: func(cCtx *cli.Context, b bool) error {
					file := cCtx.Args().Get(0)
					countString, _ := lineCount(cCtx, file)
					fmt.Println((countString))
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "words",
				Aliases: []string{"w"},
				Usage:   "displays the number of words in the file",
				Action: func(cCtx *cli.Context, b bool) error {
					file := cCtx.Args().Get(0)
					countString, _ := wordCount(cCtx, file)
					fmt.Println((countString))
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "bytes",
				Aliases: []string{"c"},
				Usage:   "displays the count of bytes in the file",
				Action: func(cCtx *cli.Context, b bool) error {
					file := cCtx.Args().Get(0)
					countString, _ :=  byteCount(cCtx, file)
					fmt.Println((countString))
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "char",
				Aliases: []string{"m"},
				Usage:   "displays count of characters from a file",
				Action: func(cCtx *cli.Context, b bool) error {
					file := cCtx.Args().Get(0)
					countString, _ := byteCount(cCtx, file)
					fmt.Println((countString))
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.NumFlags() == 0 {
				file := cCtx.Args().Get(0)
				allCount := getAllCount(cCtx, file)
				fmt.Println(allCount)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
