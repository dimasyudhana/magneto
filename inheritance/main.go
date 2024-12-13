package main

import (
	"fmt"
	"log"
)

type Human struct {
	Firstname string
	Lastname  string
}

func (t Human) Walk() {
	fmt.Print("I'm walk ...\n")
	fmt.Printf("saldo: %s\n", t.Firstname)
}

type SchoolTeacher struct {
	Human Human
}

type UniversityTeacher struct {
	Human       Human
	departement string
	speciality  string
}

func main() {

	jhon := UniversityTeacher{
		Human: Human{
			Firstname: "JP",
			Lastname:  "Morgan",
		},
		departement: "Sci",
		speciality:  "Physics",
	}

	jane := SchoolTeacher{
		Human: Human{
			Firstname: "JP",
			Lastname:  "Gaby",
		},
	}

	log.Printf("%+v", jhon)
	log.Printf("%+v", jane)

	jhon.Human.Walk()
	jane.Human.Walk()

}
