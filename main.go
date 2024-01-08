package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("El pais brazil tiene %d viajes\n", total)

	totalMornings, errDos := tickets.GetMornings("Morning")
	if errDos != nil {
		println(err.Error())
		return
	}
	fmt.Printf("Hubo %d viajes en la mañana\n", totalMornings)

	destino := "Brazil"
	averageDestination, errTres := tickets.AverageDestination(destino)
	if errTres != nil {
		println(err.Error())
		return
	}
	fmt.Printf("El porcentaje de viejes hacia %s, es de %0.2f por ciento\n", destino, averageDestination)
}
