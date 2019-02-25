package main

import "fmt"

func main() {
	str1 := "-1234567"
	fmt.Println(myAtoi(str1))
	str2 := "4193 with words"
	fmt.Println(myAtoi(str2))
	str3 := "-91283472332"
	fmt.Println(myAtoi(str3))
}
