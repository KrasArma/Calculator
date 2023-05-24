package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	en()
}

func en() {
	x, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lst := strings.Fields(x)
	chek1 := [4]string{"+", "-", "/", "*"}
	marker := 0
	if len(lst) == 3 {
		for i := range chek1 {
			if chek1[i] == lst[1] {
				marker = 1
				break
			}
		}
	}

	if marker != 1 {
		fmt.Println("Ошибка. Формат математической операции не удовлетворяет заданию.")
	} else {
		control(lst)
	}
}

func control(lst []string) {
	marker := [3]string{"", "", ""}
	transl := map[string]string{
		"1":  "I",
		"2":  "II",
		"3":  "III",
		"4":  "IV",
		"5":  "V",
		"6":  "VI",
		"7":  "VII",
		"8":  "VIII",
		"9":  "IX",
		"10": "X",
	}
	for i := 0; i < 3; i += 2 {
		for key, value := range transl {
			if lst[i] == key {
				marker[i] = "a"
				break
			} else if lst[i] == value {
				marker[i] = "r"
				lst[i] = key
				break
			}
		}
	}

	if marker[0] == "" || marker[2] == "" {
		fmt.Println("Ошибка. Формат математической операции не удовлетворяет заданию.")
	} else if marker[0] != marker[2] {
		fmt.Println("Ошибка. Используются одновременно разные системы счисления.")
	} else {
		if marker[0] == "r" {
			s := a_operation(lst)
			if s <= 0 {
				fmt.Println("Ошибка. В римской системе нет отрицательных чисел.")
			} else {
				rtoa(s)
			}
		} else {
			s := a_operation(lst)
			fmt.Println(s)
		}
	}
}

func a_operation(lst []string) int {
	x, _ := strconv.Atoi(lst[0])
	y, _ := strconv.Atoi(lst[2])
	if lst[1] == "+" {
		return x + y
	} else if lst[1] == "-" {
		return x - y
	} else if lst[1] == "/" {
		return x / y
	} else {
		return x * y
	}
}

func rtoa(num int) {
	ones := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hunds := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thous := [5]string{"", "M", "MM", "MMM", "MMMM"}

	x1 := thous[num/1000]
	x2 := hunds[num/100%10]
	x3 := tens[num/10%10]
	x4 := ones[num%10]

	fmt.Println(x1 + x2 + x3 + x4)
}
