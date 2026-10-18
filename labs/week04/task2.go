package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

func (p Person) Contact() string {
	return p.Name + " <" + p.Email + ">"
}

type UStudent struct {
	Person
	GPA float64
}

type Teacher struct {
	Person
	Department string
}

func (t Teacher) Contact() string {
	return t.Name + " <" + t.Email + "> " + t.Department
}

type Passport struct {
	ID string
}

type License struct {
	ID string
}

type Citizen struct {
	Passport
	License
}

func task2() {
	s := UStudent{Person{"Ali", "ali@kbtu.kz"}, 3.7}
	t := Teacher{Person{"Bakytzhan", "b@kbtu.kz"}, "Computer Science"}
	fmt.Println(s.Contact())
	fmt.Println(t.Contact())
	fmt.Println(t.Person.Contact())

	c := Citizen{Passport{"P-1"}, License{"L-2"}}
	fmt.Println(c.Passport.ID)
	fmt.Println(c.License.ID)
}
