package main

import "fmt"

func main() {
	fmt.Println(calendar(1))
}

func calendar(dayOfWeek int) string {
	switch dayOfWeek {
	case 1:
		return "Сегодня пн"
	case 2:
		return "Сегодня вт"
	case 3:
		return "Сегодня ср"
	case 4:
		return "Сегодня чт"
	case 5:
		return "Сегодня пт"
	case 6:
		return "Сегодня сб"
	case 7:
		return "Сегодня вс"
	default:
		return "Такого дня нет"
	}
}
