package tickets

type CountByCountry struct {
}

func (task CountByCountry) OperateLine(line []string, bag *Bag) {
	ticketMap, ok := (bag.Object).(map[string]int) // REVISAR LÓGICA DE PUNTEROS
	if ok {
		StoreLineInMap(line, &ticketMap)
	} else {
		return // error
	}
}

func StoreLineInMap(line []string, ticketsMap *map[string]int) {
	destination := line[3]

	if value, ok := (*ticketsMap)[destination]; ok {
		(*ticketsMap)[destination] = value + 1
	} else {
		(*ticketsMap)[destination] = 1
	}
}
