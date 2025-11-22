package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cnvFlags = flag.NewFlagSet("cnv", flag.ExitOnError)
	width    = cnvFlags.Int64("width", 1024, "")
	thumb    = cnvFlags.Bool("thumb", false, "")
	filePath = cnvFlags.String("filePath", "", "")

	filterFlags = flag.NewFlagSet("filter", flag.ExitOnError)
	isGray      = filterFlags.Bool("gray", false, "")
	isSepia     = filterFlags.Bool("sepia", false, "")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("set or get subcomand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cnv":
		cnvFlags.Parse(os.Args[2:])
	case "filter":
		filterFlags.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if cnvFlags.Parsed() {
		fmt.Println(*width)
		fmt.Println(*thumb)
		fmt.Println(*filePath)
	}

	if filterFlags.Parsed() {
		fmt.Println(*isGray)
		fmt.Println(*isSepia)
	}

	//for i, v := range flag.Args() {
	//fmt.Println(i, "-", v)
	//}
}
