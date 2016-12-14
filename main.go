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
				cli.BoolFlag{
					Name:  "d",
					Usage: "Run service in detached mode",
				},
			},
		},
	}
}

func startService(c *cli.Context) error {
    // define required variables
    var err error
	var app      = c.Args().Get(0)
    var detached = c.Bool("d")

    // build command
    cmd := exec.Command("./" + app)

    // if detached process is requested
    // else output the command to stdout
    fmt.Println("Starting " + app + " service...")
    if detached {
        err = cmd.Start()
    } else {
        _,err = cmd.Output()
    }

    fmt.Println(cmd.Stdout)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
