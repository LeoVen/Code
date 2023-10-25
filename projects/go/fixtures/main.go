package main

import (
	"fixtures_example/grades"
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Open("grades.csv")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
	}
	grades, err := grades.NewGradebook(csvFile)
	fmt.Printf("%+v\n", grades.FindByStudent("Jane"))
}
