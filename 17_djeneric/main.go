package main

import fmtprinter "fmtPrinter/fmtPrinter"

// базовая типизация (в Go это делается черз interface)
type Number interface {
	int64 | float64
}

type ConsolePrinter interface {
	Println(elem any)
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1, 2, 3, 4, 5, 6, 7}
	// c := []string{"1", "2", "3"}

	printElements(sum(a), &fmtprinter.FmtPrinter{})
	printElements(sum(b), &fmtprinter.FmtPrinter{})
}

// дженерик функция
func sum[T Number](arr []T) T {
	var res T
	for _, num := range arr {
		res += num
	}
	return res
}

// функция вывод с внедрением зависимостей fmt
func printElements(elm any, printer ConsolePrinter) {
	printer.Println(elm)
}
