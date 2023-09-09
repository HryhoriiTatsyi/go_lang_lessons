package main

import (
	"fmt"
	"sort"

	"./zoo"
)

func main() {
	fmt.Printf("Welcome to our zoo, for you available few types of tickets: \n \n")
	ticket_buy()
	fmt.Printf("\n")
	animals_in_zoo()
}

func ticket_buy() {
	tickets := []zoo.Ticket{
		{
			Type:  "Simple",
			Desc:  "this ticket allow you to visit non famosable animals",
			Price: 25,
		},
		{
			Type:  "Lux",
			Desc:  "this ticket allow you to do what you wanna do",
			Price: 50,
		},
	}
	for _, tickets_type := range tickets {
		fmt.Printf("Type: %s, Description: %s, Price: %d\n", tickets_type.Type, tickets_type.Desc, tickets_type.Price)
	}
}

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
			Age:       15,
			Famosable: false,
		},
		{
			Name:      "Crocodile",
			Continent: "Afrika",
			Behaivor:  "agressive",
			Size:      "large",
			Age:       150,
			Famosable: true,
		},
	}

	sort.Slice(animals, func(i, j int) bool {
		return animals[i].Famosable && !animals[j].Famosable
	})

	fmt.Print("Here the list of famous animals:\n")
	for _, famous_animals_in_zoo := range animals {
		if famous_animals_in_zoo.Famosable {
			fmt.Printf("Name: %s, Continent: %s, Behaivor: %s, Size: %s, Age: %d\n", famous_animals_in_zoo.Name, famous_animals_in_zoo.Continent, famous_animals_in_zoo.Behaivor, famous_animals_in_zoo.Size, famous_animals_in_zoo.Age)
		}
	}

	fmt.Print("\nHere the list of unfamous animals:\n")
	for _, unfamous_animals_in_zoo := range animals {
		if unfamous_animals_in_zoo.Famosable {
		} else {
			fmt.Printf("Name: %s, Continent: %s, Behaivor: %s, Size: %s, Age: %d\n", unfamous_animals_in_zoo.Name, unfamous_animals_in_zoo.Continent, unfamous_animals_in_zoo.Behaivor, unfamous_animals_in_zoo.Size, unfamous_animals_in_zoo.Age)
		}
	}

}
