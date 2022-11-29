package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Математические функции
func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
}

func checkNumberSize(firstNumber, secondNumber int) bool {
	if firstNumber > 10 || secondNumber > 10 || secondNumber == 0 || firstNumber == 0 {
		return false
	} else {
		return true
	}
}

//Функции конвертации
func convertToRoman(num int) string {
	numbers := [15]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 3, 2, 1}
	symbols := [15]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "III", "II", "I"}
	var resultString string = ""

	for num > 0 {
		for i := 0; i < len(numbers)-1; i++ {
			wholeDivision := num / numbers[i]
			if wholeDivision > 0 {
				num -= (wholeDivision * numbers[i])
				for k := 0; k < wholeDivision; k++ {
					resultString += symbols[i]
				}

			}
		}
	}
	return resultString
}

func convertToArabic(num string) int {

	romanSymbolsMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lastRomanDigit := string(num[len(num)-1])                 //последняя цифра в римском числе
	var resultArabicNum int = romanSymbolsMap[lastRomanDigit] //результирующее число
	for i := len(num) - 2; i >= 0; i-- {
		if romanSymbolsMap[string(num[i])] < romanSymbolsMap[string(num[i+1])] {
			resultArabicNum -= romanSymbolsMap[string(num[i])]
		} else {
			resultArabicNum += romanSymbolsMap[string(num[i])]
		}
	}
	return resultArabicNum
}

func main() {
	var romanExample bool = false //переменная определяющая в какой системе цифр пример
	var firstNum, secondNum int = 0, 0
	var resultNumber int = 0
	operators := "+-*/"
	numbers := "0123456789"
	romanSymbols := "IVXLCDM"
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter an example of 2 Arabic or Roman numerals")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		operatorIndex := strings.IndexAny(text, operators)
		hasArabic := strings.ContainsAny(text, numbers)
		hasRoman := strings.ContainsAny(text, romanSymbols)

		//проверка на наличие одновременно разных цифр
		if hasRoman && hasArabic {
			println("[ERROR] You cannot use Roman and Arabic numerals at the same time")
			return
		} else if hasRoman {
			//определение того, что пример состоит только из римских цифр
			romanExample = true
		}

		if operatorIndex == -1 {
			println("[ERROR] No operator")
			return
		}

		operator := string(text[operatorIndex])

		separatedText := strings.Split(text, operator)
		if len(separatedText) < 2 || len(separatedText) > 2 {
			println("[ERROR] Too many numbers or just one number")
			return
		}
		//получение цифр, в зависимости от типа цифр
		if romanExample {
			firstNum = convertToArabic(separatedText[0])  //получение первой цифры
			secondNum = convertToArabic(separatedText[1]) //получение второй цифры
		} else {
			firstNum, _ = strconv.Atoi(separatedText[0])  //получение первой цифры
			secondNum, _ = strconv.Atoi(separatedText[1]) //получение второй цифры
		}

		if checkNumberSize(firstNum, secondNum) == false {
			println("[ERROR] The numbers should be between 1 and 10")
			return
		}

		switch operator {
		case "+":
			resultNumber = add(firstNum, secondNum)
		case "-":
			resultNumber = subtract(firstNum, secondNum)
		case "*":
			resultNumber = multiply(firstNum, secondNum)
		case "/":
			resultNumber = divide(firstNum, secondNum)
		}
		//проверка на отрицательный результат для римских цифр
		if resultNumber < 1 && romanExample {
			println("[ERROR] Not a positive result")
			return
		} else if romanExample {
			println("Result : ", convertToRoman(resultNumber))
		} else {
			println("Result : ", resultNumber)
			println("__________________________________________") //Строка для отделения примеров друг от друга
		}
	}
}
