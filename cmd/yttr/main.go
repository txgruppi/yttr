package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/txgruppi/yttr"
)

const (
	ECUknownErr   = 127
	ECNoApiPair   = 126
	ECNoFile      = 125
	ECYttrFileErr = 124
	ECUploadErr   = 123
)

func main() {
	app := cli.NewApp()
	app.Name = yttr.Name
	app.Version = yttr.Version
	app.Usage = "Upload file to yttr.co from your command line"
	app.Authors = []cli.Author{
		{Name: "Tarcisio Gruppi", Email: "txgruppi@gmail.com"},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "API key",
			EnvVar: "YTTR_API_KEY",
		},
		cli.StringFlag{
			Name:   "secret, s",
			Usage:  "API secret",
			EnvVar: "YTTR_API_SECRET",
		},
		cli.BoolFlag{
			Name:  "downloadOnly, d",
			Usage: "Make the file available only for download. Default false",
		},
		cli.IntFlag{
			Name:  "expire, e",
			Usage: "After how many days the file must expire. Can be 1, 4 or 12. Default 1",
			Value: 1,
		},
	}
	app.Action = func(c *cli.Context) {
		var exitCode int
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				os.Exit(ECUknownErr)
			}

			if exitCode != 0 {
				os.Exit(exitCode)
			}
		}()

		exitCode = run(c)
	}
	app.Run(os.Args)
}

func run(c *cli.Context) int {
	key := c.String("key")
	secret := c.String("secret")
	downloadOnly := c.Bool("downloadOnly")
	expire := c.Int("expire")

	if len(key) != 40 || len(secret) != 40 {
		fmt.Fprintf(os.Stderr, "A API key/secret pair is required.\n")
		return ECNoApiPair
	}

	if len(c.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "A file is required.\n")
		return ECNoFile
	}

	yFile, err := yttr.NewFileFromPath(c.Args()[0], expire, downloadOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create a yttr.File: %s\n", err.Error())
		return ECYttrFileErr
	}

	req := yttr.NewUploadRequest(yFile)
	api := yttr.NewAPI(key, secret)

	res, err := api.Upload(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't upload file: %s\n", err.Error())
		switch err.(type) {
		case yttr.Error:
			return err.(yttr.Error).Type()
		default:
			return ECUknownErr
		}
	}

	fmt.Println(res.URL().String())

	return 0
}
