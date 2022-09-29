package main

import (
	"exam"
	"fmt"
)

func main() {
	fmt.Println(exam.PrintElements([]int{1, 4}, []string{"Choice", "Hello", "Solid", "Curtain", "There", "Forward"}))
	fmt.Println(exam.PrintElements([]int{2, 0, 7}, []string{"Kenobi", "Unity", "General", "Therapist"}))
}
