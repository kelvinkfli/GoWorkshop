package main

import "fmt"

type workshop struct {
	Location  string
	Attendees []string
	HasFood   bool
}

func main() {
	var w workshop

	w.Location = "CrowdRiff"
	w.Attendees = []string{"really", "cool", "people"}
	w.HasFood = true

	w2 := workshop{
		Location:  "116 Spadina",
		Attendees: []string{"More", "cool", "people"},
		HasFood:   true,
	}

	fmt.Printf("%+v \n", w)
	fmt.Printf("%+v", w2)
}
