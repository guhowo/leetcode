package main

import (
	"strings"
	"math"
)

func myAtoi(str string) int {
    return convert(clean(str))
}

func clean(str string) (int, string) {
    str = strings.TrimSpace(str)
	if str == "" {
		return 0, ""
	}
    sign := 1
    switch str[0]{
        case '0','1','2','3','4','5','6','7','8','9':
            sign, str = 1, str[:]
        case '+':
            sign, str = 1, str[1:]
        case '-':
            sign, str = -1, str[1:]
        default:
            return 0, ""
    }

    for i, v := range str {
        if v <'0' || v >'9' {
			return sign, str[:i]
        }
    }
    return sign, str
}

func convert(sign int, str string) int {
    num := 0
    for _, v := range str {
        num = num*10 + int(v-'0')
		switch {
		case sign == 1 && num > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && sign*num < math.MinInt32:
			return math.MinInt32
		}
    }
    return sign*num
}
