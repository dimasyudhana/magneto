package main

import "fmt"

type Teacher interface {
	Greeting()
}

type SchoolTeacher struct {
	id        int
	firstname string
	lastname  string
}

type UniversityTeacher struct {
	id         int
	firstname  string
	lastname   string
	department string
	specialty  string
}

func (s SchoolTeacher) Greeting() {
	fmt.Printf("%s %s\n", s.firstname, s.lastname)
	s.Helper()
}

func (s SchoolTeacher) Helper() {
	fmt.Printf("%s %s\n", s.firstname, s.lastname)
}

func (s UniversityTeacher) Greeting() {
	fmt.Printf("%s %s %s %s\n", s.firstname, s.lastname, s.department, s.specialty)
}

func main() {

	john := UniversityTeacher{
		id:         101,
		firstname:  "JP",
		lastname:   "Morgan",
		department: "sci",
		specialty:  "physics",
	}

	jane := SchoolTeacher{
		id:        102,
		firstname: "Jane",
		lastname:  "Gaby",
	}

	teachers := []Teacher{
		john,
		jane,
	}

	for _, val := range teachers {
		val.Greeting()
	}

}
