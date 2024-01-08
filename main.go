package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	destino := "Brazil"
	total, err := tickets.GetTotalTickets(destino)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("El pais %s tiene %d viajes\n", destino, total)

	totalMornings, errDos := tickets.GetMornings("Morning")
	if errDos != nil {
		println(err.Error())
		return
	}
	fmt.Printf("Hubo %d viajes en la mañana\n", totalMornings)

	averageDestination, errTres := tickets.AverageDestination(destino)
	if errTres != nil {
		println(err.Error())
		return
	}
	fmt.Printf("El porcentaje de viejes hacia %s, es de %0.2f por ciento\n", destino, averageDestination)
}
