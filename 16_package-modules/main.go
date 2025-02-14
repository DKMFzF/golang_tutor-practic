package main

import (
	"context"
	"fmt"
	"golang-ninja/basic/shape"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

func main() {
	square := shape.NewCircle(2.0)
	circle := shape.NewCircle(5.0)

	fmt.Println(square)
	fmt.Println(circle)

	printShapeArea(&square)
	printShapeArea(&circle)

	printShapePerimetr(&square)
	printShapePerimetr(&circle)

	// действия с импортируемой библиотекой
	// scheduler := scheduler.NewScheduler()

	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics sent at %s\n", time.Now().String())
}

func printShapeArea(s shape.Shape) {
	fmt.Println("Площадь ->", s.Area())
}

func printShapePerimetr(s shape.Shape) {
	fmt.Println("Периметр ->", s.Perimetr())
}
