package main

import "fmt"

func main() {
	const TotalAllowedNiggas int = 5
	var remainingNiggas int = 5
	var NiggaName string
	var NiggaCount int

	fmt.Print("Please Enter your name:")
	fmt.Scan(&NiggaName)
	fmt.Printf("Hey %v Welcome to The GO - Application\n", NiggaName)
	fmt.Printf("We have Total: %v And Remaining: %v in my Basement\n", TotalAllowedNiggas, remainingNiggas)
	fmt.Print("Please Enter the number of Niggas to Book:")
	fmt.Scan(&NiggaCount)
	fmt.Printf("%v Booked %v Niggers for their Basement", NiggaName, NiggaCount)
}
