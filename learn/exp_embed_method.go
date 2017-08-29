package main

import (
	"fmt"
	"log"
)

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Println("user ", u)
	return nil
}

type Admin struct {
	User
	Level string
}

func (u *Admin) Notify() error {
	log.Println("admin ", u)
	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	fmt.Println("testing embedded receiver")
	bill := &User{"Bill", "bill@google.com"}
	bill.Notify()
	SendNotification(bill)

	admin1 := &Admin{User{"Jill", "Jill@google.com"}, "1"}
	admin1.Notify()
	admin1.User.Notify()
	log.Println("done")
}
