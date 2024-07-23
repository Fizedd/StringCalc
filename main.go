package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")
	expr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	expr = strings.TrimSpace(expr)

	// Проверка длины входного выражения
	if len(expr) > 10 {
		panic("входное выражение должно содержать не более 10 символов")
	}

	switch {
	case strings.Contains(expr, " + "):
		parts := strings.Split(expr, " + ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		str1 := strings.Trim(parts[0], "\"")
		str2 := strings.Trim(parts[1], "\"")
		result := str1 + str2
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(expr, " - "):
		parts := strings.Split(expr, " - ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		str1 := strings.Trim(parts[0], "\"")
		str2 := strings.Trim(parts[1], "\"")
		result := strings.Replace(str1, str2, "", 1)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(expr, " * "):
		parts := strings.Split(expr, " * ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		str := strings.Trim(parts[0], "\"")
		num, err := strconv.Atoi(parts[1])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		result := strings.Repeat(str, num)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(expr, " / "):
		parts := strings.Split(expr, " / ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		str := strings.Trim(parts[0], "\"")
		num, err := strconv.Atoi(parts[1])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		partLength := len(str) / num
		if partLength == 0 {
			fmt.Println("")
		} else {
			result := str[:partLength]
			if len(result) > 40 {
				result = result[:40] + "..."
			}
			fmt.Println(result)
		}

	default:
		panic("неверное выражение")
	}
}
