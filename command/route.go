package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/crerwin/septa/septa"
)

type RouteCommand struct {
}

func (c *RouteCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println("please specify a route number")
		return 1
	}
	var routeNumber = args[0]
	url := fmt.Sprintf("http://www3.septa.org/beta/TransitView/" + routeNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return 1
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return 1
	}
	defer resp.Body.Close()

	var record septa.TransitViewResult

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Route", routeNumber, "-", len(record.Bus), "buses:")
	for _, bus := range record.Bus {
		fmt.Printf("Bus No. %v %v toward %v\n", bus.VehicleID, bus.Direction,
			bus.Destination)
	}
	return 0
}

func (c *RouteCommand) Help() string {
	return "route help"
}

func (c *RouteCommand) Synopsis() string {
	return "route synopsis"
}
