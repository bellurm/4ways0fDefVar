package main

import "fmt"

func print4way() {
	var studentName, grade, isPassed = "John Doe", 77, true
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

	/* studentName, grade, isPassed := "John Doe", 77, true
	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
	*/

	/*
		var (
			studentName string = "John Doe"
			grade       int    = 77
			isPassed    bool   = true
		)

		fmt.Println(studentName)
		fmt.Println(grade)
		fmt.Println(isPassed)
	*/

	/*
		var studentName string = "John Doe"
		var grade int = 77
		var isPassed bool = true

		fmt.Println(studentName)
		fmt.Println(grade)
		fmt.Println(isPassed)
	*/
	/*
		studentName := "John Doe"
		grade := 77
		isPassed := true

		fmt.Println(studentName)
		fmt.Println(grade)
		fmt.Println(isPassed)
	*/
}
