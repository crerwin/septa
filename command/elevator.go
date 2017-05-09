package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ElevatorCommand struct{}

func (c *ElevatorCommand) Run(args []string) int {

	url := fmt.Sprintf("http://www3.septa.org/hackathon/elevator/")

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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	bs := string(body)

	fmt.Println("Elevator status:")
	fmt.Println(bs)

	return 0
}

func (c *ElevatorCommand) Help() string {
	return "elevator status help"
}

func (c *ElevatorCommand) Synopsis() string {
	return "show elevator status alerts"
}
