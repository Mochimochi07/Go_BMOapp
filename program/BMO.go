package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Booting...\n")
	time.Sleep(2 * time.Second)
	fmt.Println("I'm awake!\n")
	fmt.Println("Who wants to play videogames?\n")
	time.Sleep(2 * time.Second) 

	for {
		fmt.Println("Title: BMO\n")
		fmt.Println("1. Play games\n")
		fmt.Println("2. Play music\n")
		fmt.Println("3. Exit\n")

		var input int
		fmt.Println("\n")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Println("Playing games...\n")
		case 2:
			fmt.Println("Playing music...\n")
		case 3:
			fmt.Println("Exiting...\n")
			return
		default:
			fmt.Println("Invalid option, please try again.\n")
		}
	}
}
