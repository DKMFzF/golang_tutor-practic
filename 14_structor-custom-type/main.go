package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type Weight int

func (w Weight) isFatHuman() bool {
	return w >= 100
}

type Height int

func (h Height) isLongHuman() bool {
	return h > 180
}

// структура которую можно использовать везде
type User struct {
	name   string
	sex    string
	age    Age
	weight Weight
	height Height
}

// конструктор для User
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age), // приводим тип к int
		weight: Weight(weight),
		height: Height(height),
	}
}

// метод для User назвается ресивер
// ресивер может быть как по значениею так и ссылкой
// Когда мы передаём значение, то мы копируем
// Когда мы передаём ссылку то копируется адресс
func (u *User) getUserName() string {
	return u.name
}

func (u *User) setUserName(name string) {
	u.name = name
}

func (u *User) printUserInfo() {
	fmt.Println(u.name, u.sex, u.age, u.weight, u.height)
}

// структура DataBase
type Database struct {
	m map[string]string
}

func NewDataBase() *Database {
	return &Database{
		m: make(map[string]string),
	}
}

func main() {
	// статичная структура которая будет доступна только в функции main()
	// user := struct {
	// 	name   string
	// 	age    int
	// 	sex    string
	// 	weight int
	// 	height int
	// }{"Kirill", 23, "male", 75, 185}

	user1 := User{"Kira", "female", 23, 64, 177}
	user2 := User{"Kirill", "male", 23, 75, 185}
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user1)

	// создание объектов user через конструктор
	user3 := NewUser("Max", "male", 23, 75, 185)
	fmt.Printf("%+v\n", user3)

	// сеттер и геттер name в User
	fmt.Println(user1.getUserName())
	user1.setUserName("Alina")
	fmt.Println(user1.getUserName())

	// DataBase
	db := NewDataBase()
	db.m["Kirill"] = user2.sex
	fmt.Println(db.m["Kirill"])

	// проверки
	if user1.age.isAdult() {
		fmt.Println("user1 is adult")
	} else {
		fmt.Println("user1 is not adult")
	}

	if user1.weight.isFatHuman() {
		fmt.Println("user1 is fat")
	} else {
		fmt.Println("user1 is not fat")
	}

	if user1.height.isLongHuman() {
		fmt.Println("user1 is long")
	} else {
		fmt.Println("user1 is not long")
	}
}
