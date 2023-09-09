package main

import (
	"fmt"

	"./zoo"
)

func main() {
	animals_in_zoo()
}

// func ticket_buy() {
// 	tickets := []zoo.Ticket{
// 		{
// 			Type:  "Simple",
// 			Price: 25,
// 		},
// 		{
// 			Type:  "Lux",
// 			Price: 50,
// 		},
// 	}
// }

func animals_in_zoo() {
	animals := []zoo.Animal{
		{
			Name:      "Zibra",
			Continent: "Afrika",
			Behaivor:  "calm",
			Size:      "middle",
			Age:       25,
			Famosable: true,
		},
		{
			Name:      "Lion",
			Continent: "Afrika",
			Behaivor:  "agressive",
			Size:      "middle",
			Age:       25,
			Famosable: false,
		},
	}

	for _, animals_in_zoo := range animals {
		fmt.Printf("Name: %s %s, Age: %d\n", animals_in_zoo.Name, animals_in_zoo.Continent, animals_in_zoo.Age)
	}
}
