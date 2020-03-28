package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/comfortablynumb/rapidito/generator/types/goginrestapi"
	rapidito2 "github.com/comfortablynumb/rapidito/rapidito"
	"github.com/urfave/cli/v2"
)

func main() {
	rapidito := rapidito2.NewRapidito()

	rapidito.RegisterGenerator(goginrestapi.NewGoGinRestApiGenerator())

	executable, err := os.Executable()

	if err != nil {
		rapidito.HandleError(err, "Could not autodetect the path of the binary. Please, set the path to the configuration helper manually.")
	}

	defaultConfigFile := fmt.Sprintf("%s/.rapidito.yaml", filepath.Dir(executable))

	app := &cli.App{
		Name:  "Rapidito",
		Usage: "@TODO: ADD USAGE EXPLANATION!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Value:    defaultConfigFile,
				Usage:    "Configuration helper",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return rapidito.Generate(c.String("file"))
		},
	}

	err = app.Run(os.Args)

	rapidito.HandleIfError(err, "There was an error while executing Rapidito!")
}
