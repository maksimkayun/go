package main

import (
	"fmt"
	"os"
)

func main() {
	//reverseNum()
	//threeNumbers()
	dateTime(32127)
}

func reverseNum() {
	var num, rev int

	fmt.Fscanln(os.Stdin, &num)

	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num /= 10
	}

	fmt.Println(rev)
}

func threeNumbers() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	maxSide := max(a, b, c)
	fmt.Println(maxSide)
	var result string
	if a == maxSide {
		if a*a == b*b+c*c {
			result = "Прямоугольный"
		} else {
			result = "Непрямоугольный"
		}

	} else if b == maxSide {
		if b*b == a*a+c*c {
			result = "Прямоугольный"
		} else {
			result = "Непрямоугольный"
		}
	} else {
		if c*c == a*a+b*b {
			result = "Прямоугольный"
		} else {
			result = "Непрямоугольный"
		}
	}

	fmt.Println(result)
}

func dateTime(input int) {
	//var input int
	//fmt.Scanf("%d", &input)
	h := int(input / 60 / 60)
	m := int(input/60) - h*60

	fmt.Printf("It is %d hours %d minutes", h, m)
}
