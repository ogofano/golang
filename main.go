package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func isLeapYear(year int) bool {
	if year % 400 == 0 {
		return true
	} else if year % 4 == 0 && year % 100 != 0 {
		return true
	}
	return false
}

func calculator(expression []string) []string {
	
	start, end := 0, 0
	counter := 0
	for i, item := range expression {
		if item == "(" {

			start = i
			counter++
			for j := i + 1; j != len(expression); j++ {
				if expression[j] == "(" {
					counter++
				} else if expression[j] == ")" {
					counter--
					if counter == 0 {
						end = j
						
						return calculator(append(append(expression[:start], calculator(expression[start + 1:end])...), expression[end + 1:]...))
					}
				}
			}
		}
	}

	for _, item := range []string{"*", "/", "+", "-"} {
		result := 0
		operand := slices.Index(expression, item)
		for operand != -1 {
			num1, _ := strconv.Atoi(expression[operand - 1])
			num2, _ := strconv.Atoi(expression[operand + 1])
			switch item {
				case "*":
					result = num1 * num2
				case "/":
					result = num1 / num2
				case "+":
					result = num1 + num2
				case "-":
					result = num1 - num2
			}
			expression = append(append(expression[:operand - 1], strconv.Itoa(result)), expression[operand + 2:]...)
			operand = slices.Index(expression, item)
		}

	}

	
	return expression
}

func main() {
	// Решение для високосного года
	// var year int
	// fmt.Scan(&year)
	// fmt.Println(isLeapYear(year))

	// Решение для калькулятора
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// pattern := regexp.MustCompile("[()\\/*\\-+0-9]")
	// fmt.Println(calculator(pattern.FindAllString(input, -1))[0])
}