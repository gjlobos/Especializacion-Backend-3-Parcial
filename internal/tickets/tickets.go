package tickets

import (
	"fmt"
	"os"
	"strings"
)

// Defino una estructura para los tickets
type Ticket struct {
	Id                 string
	Name               string
	Email              string
	DestinationCountry string
	FlightHour         string
	Price              string
}

// Requerimiento 1
// Una función que calcule cuántas personas viajan a un país determinado.
func GetTotalTickets(DestinationCountry string) (int, []Ticket, error) {
	file, err := os.ReadFile("./tickets.csv")

	if err != nil {
		fmt.Println("Se produjo un error al leer el archivo", err)
	}

	tickets := strings.Split(string(file), "\n")

	countTickets := 0
	var ticketsResults []Ticket

	for i := 0; i < len(tickets); i++ {
		eachTicket := strings.Split(tickets[i], ",")

		if eachTicket[3] == DestinationCountry {
			countTickets++

			ticket := Ticket{
				Id:                 eachTicket[0],
				Name:               eachTicket[1],
				Email:              eachTicket[2],
				DestinationCountry: eachTicket[3],
				FlightHour:         eachTicket[4],
				Price:              eachTicket[5],
			}

			ticketsResults = append(ticketsResults, ticket)
		}
	}

	return countTickets, ticketsResults, nil
}

// Requerimiento 2
// Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6),
// mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
//func GetMornings(time string) (int, error) {}

// Requerimiento 3
// Calcular el promedio de personas que viajan a un país determinado en un día.
// func AverageDestination(destination string, total int) (int, error) {}
