package tickets

type CountByAverage struct {
}

func (task CountByAverage) OperateLine(line []string, bag *Bag) {
	ticketMap, ok := (bag.Object).(map[string]int) // REVISAR LÓGICA DE PUNTEROS
	if ok {
		countAvergaAndTtotal(line, &ticketMap)
	} else {
		return // error
	}
}

func countAvergaAndTtotal(line []string, ticketsMap *map[string]int) {
	destination := line[3]

	(*ticketsMap)["Total"] += 1

	if value, ok := (*ticketsMap)[destination]; ok {
		(*ticketsMap)[destination] = value + 1
	} // solo suma si el destintation fue declarado anteriormente
}
