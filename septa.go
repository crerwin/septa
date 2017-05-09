package main

import (
	"log"
	"os"

	"github.com/crerwin/septa/command"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("septa", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"route": func() (cli.Command, error) {
			return &command.RouteCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
