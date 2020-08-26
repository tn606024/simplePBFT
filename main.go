package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			PBFTCommand,
		},
	}


	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("%s",err)
	}
}
