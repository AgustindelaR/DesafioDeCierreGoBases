package tickets

import (
	"encoding/csv"
	"os"
)

func CsvReader(path string, task Job, bag *Bag) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		/*
			ticket := Ticket{
				Mail:        record[0],
				Name:        record[1],
				Destination: record[2],
				...
			}*/

		err = task.OperateLine(record, bag)
	}
	return err
}
