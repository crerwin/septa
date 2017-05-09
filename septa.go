package main

import (
	"log"
	"os"

	"github.com/crerwin/septa/command"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("septa", Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"route": func() (cli.Command, error) {
			return &command.RouteCommand{}, nil
		},
		"elevator": func() (cli.Command, error) {
			return &command.ElevatorCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
