package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/crerwin/septa/septa"
)

func main() {
	fmt.Println("SEPTA command line version " + Version)
	fmt.Println("Copyright 2017 Chris Erwin")

	args := os.Args
	if len(args) == 1 {
		fmt.Println("buses")
		os.Exit(0)
	} else {
		fmt.Println(args[1])
	}

	route := args[2]
	url := fmt.Sprintf("http://www3.septa.org/beta/TransitView/" + route)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var record septa.TransitViewResult

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Route", route, "-", len(record.Bus), "buses:")
	for _, bus := range record.Bus {
		fmt.Printf("Bus No. %v %v toward %v\n", bus.VehicleID, bus.Direction,
			bus.Destination)
	}
}
