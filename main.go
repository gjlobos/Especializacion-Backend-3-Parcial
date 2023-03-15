package main

import (
	"fmt"

	"github.com/gjlobos/Especializacion-Backend-3-Parcial/internal/tickets"
)

func main() {
	totalTickets, ticketsResults, err := tickets.GetTotalTickets("Brazil")

	if err != nil {
		panic(err)
	}

	fmt.Printf("total de tickets %d \n", totalTickets)
	fmt.Printf("tickets %s \n", ticketsResults)

	//	totalTicketsByPeriod, ticketsResults, err := tickets.GetPeriods("8:00")
	//
	//	if err != nil {
	//		panic(err)
}

//	fmt.Printf("total de tickets %d \n", totalTicketsByPeriod)
//	fmt.Printf("tickets %s \n", ticketsResults)

//	averageDestination, err := tickets.AverageDestination()

//	if err != nil {
//		panic(err)
//	}

//	fmt.Println(averageDestination)
//}
