package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var convert = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

func converter(symb string) int {
	length := len(symb)
	var res int
	if length == 1 {
		res = convert[strings.ToUpper(symb)]
	} else {
		sl := strings.Split(symb, "")
		for i := 0; i < length; i++ {
			if i == 0 {
				res += convert[sl[i]]
				continue
			}
			switch {
			case convert[sl[i-1]] < convert[sl[i]]:
				res = (res - convert[sl[i-1]]) + (convert[sl[i]] - convert[sl[i-1]])
			default:
				res += convert[sl[i]]
			}
		}
	}
	return res
}
func searchMap(trg int) string {
	var res string
	for key, value := range convert {
		if value == trg {
			res = key
			break
		}
	}
	return res
}
func converterRomAr(symb int) string {
	var res string
	complete := false
	for key, value := range convert {
		if value == symb {
			res = key
			complete = true
			break
		}
	}
	if complete != true {
		var firstNum, secondNum int
		firstNum = (symb / 10) * 10
		secondNum = symb % 10
		res = searchMap(firstNum) + searchMap(secondNum)
	}

	return res
}
func calculate(a int, b int, symb string) int {
	var res int
	switch symb {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		res = 101
	}

	return res
}
func main() {

	var a_, symb_, b_ string
	var a, b int
	var romNum0, romNum1, arbNum0, arbNum1 bool
	reader := bufio.NewReader(os.Stdin)
	//input = fmt.Scan(&a_, &symb_, &b_)
	text, _ := reader.ReadString('\n')
	input := strings.Split(text, " ")
	if len(input) == 1 {
		fmt.Println("Ошибка... формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		os.Exit(0)
	}

	a_ = input[0]
	symb_ = strings.TrimSpace(input[1])
	b_ = strings.TrimSpace(input[2])

	fmt.Println(romNum0, romNum1, arbNum0, arbNum1)

	a, err := strconv.Atoi(a_)
	arbNum0 = true
	if err != nil {
		romNum0 = true
		arbNum0 = false
		a = converter(a_)
	}

	fmt.Println(romNum0, arbNum0)
	b, err = strconv.Atoi(b_)
	arbNum1 = true
	if err != nil {
		romNum1 = true
		arbNum1 = false
		b = converter(b_)
	}
	fmt.Println(romNum1, arbNum1)
	var sum = calculate(a, b, symb_)
	switch {
	case a > 10 || b > 10:
		fmt.Println("Ошибка... формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	case romNum0 != romNum1 || arbNum0 != arbNum1:
		fmt.Println("Ошибка... используются одновременно разные системы счисления")
	case sum > 100 || len(input) != 3:
		fmt.Println("Ошибка... формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	case sum < 1 && romNum0 == true && romNum1 == true:
		fmt.Println("В римской системе нет отрицательных чисел и нуля")
	case romNum0 == true && romNum1 == true:
		fmt.Println(converterRomAr(sum))
	case arbNum0 == true && arbNum1 == true:
		fmt.Println(sum)
	}
}
