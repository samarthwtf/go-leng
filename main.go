package main

import "fmt"

func main() {
	const TotalAllowedNiggas int = 5
	var remainingNiggas uint = 5
	var CustomerName string
	var SellingNiggaCount uint

	var nigga_name string
	var niggas_details [5]string

	fmt.Print("Please Enter your name: ")
	fmt.Scan(&CustomerName)
	fmt.Printf("Hey %v Welcome to Nigga Trade Centre\n", CustomerName)
	fmt.Printf("We have Total: %v Niggers\nAnd Remaining: %v Niggers", TotalAllowedNiggas, remainingNiggas)
	fmt.Print("Please Enter the number of Niggas to Book:")
	fmt.Scan(&SellingNiggaCount)
	fmt.Printf("%v Booked %v Niggers for their Basement\n", CustomerName, SellingNiggaCount)
	fmt.Printf("Please Provide the required information ")
	fmt.Print("Please Enter your nigga's Name: ")
	fmt.Scan(&nigga_name)
	niggas_details[0] = nigga_name
	remainingNiggas = remainingNiggas - SellingNiggaCount
	fmt.Printf("Remaining Niggas are : %v\n", remainingNiggas)
}
