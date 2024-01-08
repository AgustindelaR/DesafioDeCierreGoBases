package tickets

import "fmt"

type Job interface {
	OperateLine([]string, *Bag)
}

type object interface{}

type Bag struct {
	Object interface{}
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	destinationsCount := make(map[string]int) // Use map and read line by line
	bag := Bag{Object: destinationsCount}
	task := CountByCountry{}

	err := CsvReader("tickets.csv", task, &bag)
	if err != nil {
		return 0, err
	}

	return destinationsCount[destination], nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	morningsAmount := 0
	bag := Bag{&morningsAmount}

	var err error
	switch time {
	case "Madrugada":
		err = CsvReader("tickets.csv", CountByMadrugada{}, &bag)
	case "Morning":
		err = CsvReader("tickets.csv", CountByMornings{}, &bag)
	case "Afternoon":
		err = CsvReader("tickets.csv", CountByTarde{}, &bag)
	case "Noon":
		err = CsvReader("tickets.csv", CountByNoche{}, &bag)
	default:
		return morningsAmount, nil
	}

	if err != nil {
		return 0, err
	}

	return morningsAmount, nil
}

// ejemplo 3
func AverageDestination(destination string) (float64, error) {
	destinationsCount := make(map[string]int) // solo va a existir el [total] y [string]

	destinationsCount["Total"] = 0
	destinationsCount[destination] = 0

	bag := Bag{Object: destinationsCount}
	task := CountByAverage{}

	err := CsvReader("tickets.csv", task, &bag)

	fmt.Printf("Viajes de %s es %d\n", destination, destinationsCount[destination])
	return (float64(destinationsCount[destination]) / float64(destinationsCount["Total"]) * 100), err
}
