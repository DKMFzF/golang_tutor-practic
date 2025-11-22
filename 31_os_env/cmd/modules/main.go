package main

import (
	"log"
	"os"

	"github.com/caarlos0/env"
)

func main() {
	// вызов переменной
	//u := os.Getenv("USER")
	//fmt.Println(u)

	// проверка существования переменной
	//fmt.Println(os.LookupEnv("USER"))

	// все переменные окружения
	//envList := os.Environ()
	//for i := range envList {
	//fmt.Println(envList[i])
	//}

	// работа с переменным окружением через пакет caarlos0/env
	type Config struct {
		PORT string `env:"PORT" envDefault:"3000"`
	}

	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		os.Exit(1)
	}

	log.Println(cfg)
}
