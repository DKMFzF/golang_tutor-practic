package fmtprinter

import "fmt"

type FmtPrinter struct{}

func (fp *FmtPrinter) Println(elem any) {
	fmt.Println(elem)
}
