package main

import (
	"fmt"
	"mime"
	"os"
	"path"
	"strings"

	"github.com/bartmeuris/progressio"
	"github.com/codegangsta/cli"
	"github.com/txgruppi/yttr"
)

const (
	ECUknownErr           = 127
	ECNoApiPair           = 126
	ECNoFile              = 125
	ECYttrFileErr         = 124
	ECUploadErr           = 123
	ECCantOpenFileErr     = 122
	ECCantStatFileErr     = 121
	ECFileNotSupportedErr = 120
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
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "Do not print transfer information",
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
	quiet := c.Bool("quiet")

	if len(key) != 40 || len(secret) != 40 {
		fmt.Fprintf(os.Stderr, "A API key/secret pair is required.\n")
		return ECNoApiPair
	}

	if len(c.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "A file is required.\n")
		return ECNoFile
	}

	mimeType := mime.TypeByExtension(path.Ext(c.Args()[0]))
	if mimeType == "" {
		fmt.Fprintf(os.Stderr, "This file type is not supported.")
		return ECFileNotSupportedErr
	}

	if strings.IndexRune(mimeType, ';') != -1 {
		mimeType = strings.SplitN(mimeType, ";", 2)[0]
	}

	stat, err := os.Stat(c.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't stat file: %s\n", err.Error())
		return ECCantStatFileErr
	}

	file, err := os.Open(c.Args()[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file: %s\n", err.Error())
		return ECCantOpenFileErr
	}

	var yFile yttr.File
	var pCh <-chan progressio.Progress
	if quiet {
		yFile = yttr.NewFile(file, stat.Name(), mimeType, stat.Size(), expire, downloadOnly)
	} else {
		var pReader *progressio.ProgressReader
		pReader, pCh = progressio.NewProgressReader(file, stat.Size())
		yFile = yttr.NewFile(pReader, stat.Name(), mimeType, stat.Size(), expire, downloadOnly)
	}
	req := yttr.NewUploadRequest(yFile)
	api := yttr.NewAPI(key, secret)

	if !quiet {
		go reportProgress(pCh)
	}

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

	if quiet {
		fmt.Println(res.URL().String())
	} else {
		fmt.Println("\n" + res.URL().String())
	}

	return 0
}

func reportProgress(ch <-chan progressio.Progress) {
	for p := range ch {
		fmt.Fprintf(
			os.Stderr,
			"%d / %d (%.2f%%)\r",
			p.Transferred,
			p.TotalSize,
			p.Percent,
		)
	}
}
