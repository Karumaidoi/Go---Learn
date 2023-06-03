package main

import (
	"fmt"
	"strings"
)

func main() {

	// k := 75

	// if k == 72 {
	// 	fmt.Println(" This is the Number", k)
	// } else {
	// 	fmt.Println("This is not the Number")
	// }

	// for i := 0; i <= 10000; i++ {
	// 	fmt.Println("This is", i)
	// }

	// for n := 0; n <= 100; n++ {
	// 	if n%2 == 0 {
	// 		fmt.Println("This is divisible by 2", n)
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// time := time.Now().Local()
	// fmt.Println(time)

	// switch k {
	// case 72:
	// 	fmt.Println("This is the Gotcha!")

	// case 73:
	// 	fmt.Println("This is 73")

	// case 75:
	// 	fmt.Println("This is the end")

	// default:
	// 	fmt.Println("This is an integer")
	// }

	c := "Dance"
	fmt.Println(len(c))
	// d := strings.Join(strings.Split(c, "a"), "")

	e := strings.ToLower(c)

	if strings.Contains(e, "a") {
		fmt.Println("Contains the letter")
	}

	strings.ReplaceAll(e, "Happy Coding", "")

	fmt.Println(strings.Contains(e, "a"))

	fmt.Println(strings.EqualFold(e, c))

	fmt.Println((strings.Split(e, "")[1]))

	var a [5]int
	b := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	fmt.Println(len("Homey"))
	fmt.Println(append(b, 100), b[len(b)-1], b[0])

}
