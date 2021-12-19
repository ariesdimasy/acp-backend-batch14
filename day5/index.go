package main

import (
	"acp-golang/day5/answer"
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 5")

	// a := 10
	// b := 20

	// answer.Swap(&a, &b)
	// fmt.Println(a, b)

	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)

	// min, max = answer.MinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println("nilai min => ", min)
	// fmt.Println("nilai max => ", max)

	// var a = answer.Student{}

	// for i := 0; i < 6; i++ {

	// 	var name string
	// 	fmt.Print("Input " + string(i) + " Student's Name : ")
	// 	fmt.Scan(&name)
	// 	a.Name = append(a.Name, name)

	// 	var score int
	// 	fmt.Print("Input " + name + "Score : ")
	// 	fmt.Scan(&score)
	// 	a.Score = append(a.Score, score)

	// }

	// fmt.Println("\n\n Average Score Student is", a.Average())
	// scoreMax, nameMax := a.Max()
	// fmt.Println("\n\n Max Score Student is", nameMax, "(", scoreMax, ")")
	// scoreMin, nameMin := a.Min()
	// fmt.Println("\n\n Min Score Student is", nameMin, "(", scoreMin, ")")

	var menu int
	var a = answer.Person{}
	var c answer.Chiper = &a

	fmt.Println("[1] Encrypt \n[2] Decrypt \n Choose your menu ?")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Println("\n Input Student's Name: ")
		fmt.Scan(&a.Name)
		fmt.Println("\n Encode of Student's Name ", a.Name, "is : ", c.Encode())
	} else if menu == 2 {
		fmt.Println("\n Input Student's Encode Name: ")
		fmt.Scan(&a.Name)
		fmt.Println("\n Decode of Student's Name ", a.NameEncode, "is : ", c.Decode())
	} else {
		fmt.Println(" Wrong input name menu ")
	}

}
