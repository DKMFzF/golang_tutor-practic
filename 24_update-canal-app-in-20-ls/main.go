package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var acriong = []string{
	"log IN",
	"log OUT",
	"log RECORD",
	"log DELETE",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	emial string
	logs  []logItem
}

func (u User) getActiviteInfo() string {
	out := fmt.Sprintf("ID:%d; | Email:%s\nActivity Log:\n", u.id, u.emial)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at\n", i+1, item.action, item.timestamp)
	}

	return out
}

func main() {
	t := time.Now()
	// добавляем wait group
	wg := &sync.WaitGroup{} // нужен для ожидания завершения всех горутин

	users := make(chan User)
	go generationUsers(100, users)

	for user := range users {
		wg.Add(1)
		go saveUsersInfo(user, wg)
	}

	// делам блокировку для го рутины main()
	wg.Wait() // это счетчик который блокируется, если есть хоть одна задча
	// блокирует до того момента пока счетчик не станет 0 (обночляет счетчик в wg.Add())

	fmt.Println("TIME ELEPSED:", time.Since(t).String())
}

func generationUsers(count int, users chan User) {
	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			emial: fmt.Sprintf("user%d@example.com", i+1),
			logs:  generationLogs(rand.Intn(1000)), // генерирует рандомное число от 0 до 1000
		}
		// time.Sleep(time.Millisecond * 10) // допусти у нас генерация пользователей долгий процесс что бы использовать каналы
	}

	close(users) // закрытие канала
}

func generationLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    acriong[rand.Intn(len(acriong)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}

func saveUsersInfo(user User, wg *sync.WaitGroup) error {
	// time.Sleep(time.Microsecond * 10)
	fmt.Println("WRATING FILE FOR USER ID: %d\n", user.id)
	filname := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filname, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActiviteInfo())
	if err != nil {
		return err
	}

	wg.Done() // нужен для оповезения Wait Group что go-рутина выполнена

	return nil
}
