package tickets

type CountByAverage struct {
}

func (task CountByAverage) OperateLine(line []string, bag *Bag) error {
	ticketMap, ok := (bag.Object).(map[string]int)
	if ok {
		countAvergaAndTtotal(line, &ticketMap)
	} else {
		return MyError{"Error convirtiendo objeto bag a *int"} // error
	}
	return nil
}

func countAvergaAndTtotal(line []string, ticketsMap *map[string]int) {
	destination := line[3]

	(*ticketsMap)["Total"] += 1

	if value, ok := (*ticketsMap)[destination]; ok {
		(*ticketsMap)[destination] = value + 1
	} // solo suma si el destintation fue declarado anteriormente
}
