package main

import "fmt"

func main() {
	const TotalAllowedNiggas int = 5
	var remainingNiggas uint = 5
	var CustomerName string
	var SellingNiggaCount uint
	fmt.Print("Please Enter your name: ")
	fmt.Scan(&CustomerName)
	fmt.Printf("Hey %v Welcome to Nigga Trade Centre\n", CustomerName)
	fmt.Printf("We have Total: %v Niggers\nAnd Remaining: %v Niggers\n", TotalAllowedNiggas, remainingNiggas)
	fmt.Print("Please Enter the number of Niggas to Book:")
	fmt.Scan(&SellingNiggaCount)
	fmt.Printf("%v Booked %v Niggers for their Basement\n", CustomerName, SellingNiggaCount)
	fmt.Printf("Please Provide the required information ")
	var nigga_name string
	var niggas_details []string
	for {
		fmt.Print("Please Enter your nigga's Name: ")
		fmt.Scan(&nigga_name)
		niggas_details = append(niggas_details, nigga_name)
		fmt.Printf("The nigga's information provided by you is : %v\n", niggas_details)
		remainingNiggas = remainingNiggas - SellingNiggaCount
		fmt.Printf("Remaining Niggas are : %v\nThank's for trading Niggers with us", remainingNiggas)
	}
}
