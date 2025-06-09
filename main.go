package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Print("Enter exam bal: ")
	student1, _ := my_modules.InputString()

	student_grade, err := strconv.Atoi(student1)

	if err != nil {
		log.Fatal(err)
	}

	studentResult, err := my_modules.CheckExam(student_grade)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result grade: %q\n", studentResult)
}
