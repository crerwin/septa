package main

import (
	"fmt"
	"os"

	"github.com/crerwin/septa/command"
)

func main() {
	fmt.Println("SEPTA command line version " + Version)
	fmt.Println("Copyright 2017 Chris Erwin")

	args := os.Args
	if len(args) == 1 {
		fmt.Println("buses")
		os.Exit(0)
	}

	routeNumber := args[2]
	route := command.Route{
		RouteNumber: routeNumber,
	}
	route.Show()
}
