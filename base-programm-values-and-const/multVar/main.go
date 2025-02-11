package main // определяет пакет в котором храниться файл (обязательно)
import (
	"fmt"
	"reflect"
)

// импорт библиотки

// основаня функция точка входа в приложения (обязательно)
func main() {
	a, b, c := 1, 2, '3'

	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(c))
}
