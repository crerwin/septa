package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/crerwin/septa/septa"
)

type Route struct {
	RouteNumber string
}

func (r *Route) Show() {
	url := fmt.Sprintf("http://www3.septa.org/beta/TransitView/" + r.RouteNumber)

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

	fmt.Println("Route", r.RouteNumber, "-", len(record.Bus), "buses:")
	for _, bus := range record.Bus {
		fmt.Printf("Bus No. %v %v toward %v\n", bus.VehicleID, bus.Direction,
			bus.Destination)
	}
}
