package tickets

import (
	"strconv"
	"strings"
)

type CountByMornings struct {
}

func (task CountByMornings) OperateLine(line []string, bag *Bag) {
	morningsAmounts, ok := (bag.Object).(*int)
	if ok {
		CountMornings(line, morningsAmounts)
	} else {
		return // error
	}
}

func CountMornings(line []string, mournings *int) {
	str := line[4]
	parts := strings.Split(str, ":")
	hoursStr := parts[0]

	hours, _ := strconv.Atoi(hoursStr) // ESTOY IGNORANDO UN ERROR !!!!
	if hours >= 9 && hours <= 12 {
		*mournings += 1
	}
}

type CountByMadrugada struct {
}

func (task CountByMadrugada) OperateLine(line []string, bag *Bag) {
	morningsAmounts, ok := (bag.Object).(*int)
	if ok {
		CountMadrugada(line, morningsAmounts)
	} else {
		return // error
	}
}

func CountMadrugada(line []string, mournings *int) {
	str := line[4]
	parts := strings.Split(str, ":")
	hoursStr := parts[0]

	hours, _ := strconv.Atoi(hoursStr) // ESTOY IGNORANDO UN ERROR !!!!

	if hours >= 0 && hours <= 6 {
		*mournings += 1
	}
}

type CountByTarde struct {
}

func (task CountByTarde) OperateLine(line []string, bag *Bag) {
	morningsAmounts, ok := (bag.Object).(*int)
	if ok {
		CountTarde(line, morningsAmounts)
	} else {
		return // error
	}
}

func CountTarde(line []string, mournings *int) {
	str := line[4]
	parts := strings.Split(str, ":")
	hoursStr := parts[0]

	hours, _ := strconv.Atoi(hoursStr) // ESTOY IGNORANDO UN ERROR !!!!
	if hours >= 13 && hours <= 19 {
		*mournings += 1
	}
}

type CountByNoche struct {
}

func (task CountByNoche) OperateLine(line []string, bag *Bag) {
	morningsAmounts, ok := (bag.Object).(*int)
	if ok {
		CountNoche(line, morningsAmounts)
	} else {
		return // error
	}
}

func CountNoche(line []string, mournings *int) {
	str := line[4]
	parts := strings.Split(str, ":")
	hoursStr := parts[0]

	hours, _ := strconv.Atoi(hoursStr) // ESTOY IGNORANDO UN ERROR !!!!

	if hours >= 20 && hours <= 23 {
		*mournings += 1
	}
}
