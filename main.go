package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til graderFahrenheit.")
			// funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en
			// flere else-if setninger
		} else {
			fmt.Println("Venligst velg convert, average eller exit:")

		}

	}

}
