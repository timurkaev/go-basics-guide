package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	LastName    string    `json:"Фамилия"`
	Email       string    `json:"Почта"`
	dateOfBirth time.Time `json:"День рождение"`
}

// У каждого поля структуры может быть набор аннотаций, которые называются тегами (tags):
type GetUserRequest struct {
	UserId    string `json:"user_id" yaml: "user_id" format:"uuid" example:"2e263a90-b74b-11eb-8529-0242ac130003"`
	IsDeleted *bool  `json:"is_deleted,omitempty" yaml:"is_deleted"`
}

func main() {
	date := time.Date(1999, 8, 25, 0, 0, 0, 0, time.UTC)
	p := Person{
		Name:        "Ismail",
		LastName:    "Timurkaev",
		Email:       "tim@mail.ru",
		dateOfBirth: date,
	}

	jsMan, err := json.Marshal(p)

	if err != nil {
		log.Fatalln("unable marshal to json")
	}

	fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}

	fmt.Println(p)
	fmt.Println(p.Email)
	fmt.Println(p.dateOfBirth)

	c := struct{ Name string }{Name: "Ismail"}
	fmt.Println(c)
}

// Конструктор
func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
	return Person{
		Name:        name,
		Email:       email,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}
