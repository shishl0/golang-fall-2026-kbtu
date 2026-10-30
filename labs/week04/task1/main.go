package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	GPA     float64
	Courses []string
}

func addCourse(s Student, course string) {
	s.Courses = append(s.Courses, course)
}

func addCoursePtr(s *Student, course string) {
	s.Courses = append(s.Courses, course)
}

func main() {
	var s1 Student
	s2 := Student{Name: "Ali", Age: 19, GPA: 3.5}
	s3 := Student{"Bota", 20, 3.8, []string{"Go"}}
	s4 := &Student{Name: "Dias", Age: 21}

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s3)
	fmt.Printf("%+v\n", s4)

	addCourse(s2, "Algorithms")
	fmt.Printf("%+v\n", s2)

	addCoursePtr(&s2, "Algorithms")
	fmt.Printf("%+v\n", s2)
}
