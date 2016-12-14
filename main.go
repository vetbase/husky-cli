package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	// create CLI application
	app := cli.NewApp()

	// set fields
	app.Name = "Husky"
	app.Usage = "Microframework for developing services"
	app.Version = "0.1.0"
	app.Commands = registerCommands()

	// return CLI app
	app.Run(os.Args)
}

func init() {
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}
}

func registerCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "start",
			Usage:  "Start Husky service",
			Action: startService,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "d",
					Value: "detached",
					Usage: "Run service in detached mode",
				},
			},
		},
	}
}

func startService(c *cli.Context) error {
	var app = c.Args().Get(0)
	_, err := exec.Command("./" + app).Output()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
