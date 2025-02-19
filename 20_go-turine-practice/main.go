package main

import (
	"fmt"
	"math/rand"
	"os"
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
	users := generationUsers(100)
	for _, user := range users {
		saveUsersInfo(user)
	}
}

func generationUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			emial: fmt.Sprintf("user%d@example.com", i+1),
			logs:  generationLogs(rand.Intn(1000)), // генерирует рандомное число от 0 до 1000
		}
	}

	return users
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

func saveUsersInfo(user User) error {
	fmt.Println("WRATING FILE FOR USER ID: %d\n", user.id)
	filname := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filname, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActiviteInfo())
	return nil
}
