package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func upperCase(str string) string {
	return strings.ToUpper(str)
}

func ex1() {

	fmt.Println("Exercise 1:")
	reader := bufio.NewReader(os.Stdin)

	str1, _ := reader.ReadString('\n')
	str2, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Print(upperCase(str1))
	fmt.Println(upperCase(str2))
}

func ex2() {

	fmt.Println("Exercise 2:")
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')

	arr := strings.Split(str, "")
	var letter, digit int = 0, 0
	for i := 0; i < len(arr); i++ {
		if match, _ := regexp.MatchString(`[a-zA-Z]+`, arr[i]); match == true {
			letter++
		}
		if match, _ := regexp.MatchString(`\d`, arr[i]); match == true {
			digit++
		}

	}

	fmt.Printf("LETTERS %v\n", letter)
	fmt.Printf("DIGITS %v\n", digit)
}

func ex3() {

	fmt.Println("Exercise 3:")
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)

	// create a map
	b := make(map[int]int)
	for i := 1; i <= n; i++ {
		b[i] = i * i
	}

	fmt.Println(b)

}

func ex4() {

	fmt.Println("Exercise 4:")
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')

	arr := strings.Split(str, "")

	myslice := []string{"0"}
	temp := ""
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == " " {
			continue
		}
		if arr[i] != " " && arr[i] != "," && arr[i] != "." {
			temp += arr[i]
		}

		if arr[i] == "," || arr[i] == "." {
			if j == 0 {
				myslice[j] = temp
				temp = ""
				j++
			}
			if j != 0 {
				myslice = append(myslice, temp)
				temp = ""
				j++
			}
		}
	}

	fmt.Println(myslice)
}

// ex5 ------------
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func ex5() {

	fmt.Println("Exercise 5:")
	r := rect{width: 5, height: 10}
	c := circle{radius: 4}

	fmt.Print("Rectangle: ")
	measure(r)

	fmt.Print("Circle: ")
	measure(c)

}

// ----------------
func main() {
	/* 1. Write a program that accepts sequence of lines as input and prints the lines after making all characters in the sentence capitalized.

	Suppose the following input is supplied to the program:

	Hello world
	Practice makes perfect
	Then, the output should be:

	HELLO WORLD
	PRACTICE MAKES PERFECT

	*/

	ex1()

	/*	2. Write a program that accepts a sentence and calculate the number of letters and digits.

		Suppose the following input is supplied to the program:

		hello world! 123
		Then, the output should be:

		LETTERS 10
		DIGITS 3
	*/

	ex2()

	/*	3. With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

		Suppose the following input is supplied to the program: 8

		Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]
		Hints:
		Use make for the map and %v verb for the output.
	*/

	ex3()

	/*
		4. Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

		Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

		Then, the output should be: [34 67 55 33 12 98]
	*/

	ex4()

	// 5. Viết 2 hàm tính diện tích hình tròn và hình chữ nhật. Sử dụng struct khai báo, interface của golang.
	ex5()

}
