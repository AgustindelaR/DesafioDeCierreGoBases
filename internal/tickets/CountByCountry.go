package tickets

type CountByCountry struct {
}

func (task CountByCountry) OperateLine(line []string, bag *Bag) error {
	ticketMap, ok := (bag.Object).(map[string]int)
	if ok {
		StoreLineInMap(line, &ticketMap)
	} else {
		return MyError{"Error en el tipado de map dentro de bag"}
	}
	return nil // no error
}

func StoreLineInMap(line []string, ticketsMap *map[string]int) {
	destination := line[3]

	if value, ok := (*ticketsMap)[destination]; ok {
		(*ticketsMap)[destination] = value + 1
	} else {
		(*ticketsMap)[destination] = 1
	}
}
