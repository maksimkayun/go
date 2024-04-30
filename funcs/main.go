package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		var input string

		//LitBeoFLcSGBOFQxMHoluDDWcqcVgkcRoAeocXO
		fmt.Scanln(&input)

		r, _ := addStars(input)

		//L*i*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*l*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
		fmt.Println(r)
	*/

	var input int

	//9118
	fmt.Scanln(&input)

	r, _ := squaring(input)

	//811164
	fmt.Println(r)
}

func addStars(str string) (string, error) {
	return strings.Join(strings.Split(str, ""), "*"), nil
}

func squaring(num int) (int, error) {
	snum := strings.Split(strconv.Itoa(num), "")
	dgts := make([]int, len(snum))

	for i, char := range snum {
		dgt, err := strconv.Atoi(string(char))
		if err != nil {
			return -1, err
		}
		dgts[i] = dgt * dgt
	}

	rs := ""
	for _, d := range dgts {
		rs += strconv.Itoa(d)
	}

	r, err := strconv.Atoi(rs)
	if err != nil {
		return -1, err
	}

	return r, nil
}
