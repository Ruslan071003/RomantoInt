package main

import "fmt"

func romanToInt1(num1 string) int {
	dictionary := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum1 := 0
	count1 := 0

	for i := 1; i < len(num1); i++ {
		if num1[i-1:i] == "I" && (num1[i:i+1] == "V" || num1[i:i+1] == "X") {
			if num1[i:i+1] == "V" {
				sum1 += -1
			} else {
				sum1 = sum1 - 1
			}
		} else if num1[i-1:i] == "X" && (num1[i:i+1] == "L" || num1[i:i+1] == "C") {
			if num1[i:i+1] == "L" {
				sum1 += -10
			} else {
				sum1 += -10
			}
		} else if num1[i-1:i] == "C" && (num1[i:i+1] == "D" || num1[i:i+1] == "M") {
			if num1[i:i+1] == "D" {
				sum1 += -100
			} else {
				sum1 += -100
			}
		} else {
			for key, value := range dictionary {
				if num1[i-1:i] == key {
					sum1 += value
					break
				}
			}
		}
		count1 += 1
	}

	for key, value := range dictionary {
		if num1[count1:count1+1] == key {
			sum1 += value
			break
		}
	}
	return sum1
}

func romanToInt2(num2 string) int {
	dictionary := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum2 := 0
	count2 := 0

	for i := 1; i < len(num2); i++ {
		if num2[i-1:i] == "I" && (num2[i:i+1] == "V" || num2[i:i+1] == "X") {
			if num2[i:i+1] == "V" {
				sum2 += -1
			} else {
				sum2 = sum2 - 1
			}
		} else if num2[i-1:i] == "X" && (num2[i:i+1] == "L" || num2[i:i+1] == "C") {
			if num2[i:i+1] == "L" {
				sum2 += -10
			} else {
				sum2 += -10
			}
		} else if num2[i-1:i] == "C" && (num2[i:i+1] == "D" || num2[i:i+1] == "M") {
			if num2[i:i+1] == "D" {
				sum2 += -100
			} else {
				sum2 += -100
			}
		} else {
			for key, value := range dictionary {
				if num2[i-1:i] == key {
					sum2 += value
					break
				}
			}
		}
		count2 += 1
	}
	for key, value := range dictionary {
		if num2[count2:count2+1] == key {
			sum2 += value
			break
		}
	}
	return sum2
}

func main() {
	num1 := "MCMXCIV"
	num2 := "III"
	num3 := "VI"
	num4 := "IIII"
	fmt.Println(num3)
	fmt.Println(num4)
	result1 := romanToInt1(num1)
	result2 := romanToInt2(num2)
	fmt.Println(result1)
	fmt.Println(result2)
}
