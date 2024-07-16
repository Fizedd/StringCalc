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
	chtoto, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	chtoto = strings.TrimSpace(chtoto)

	switch {
	case strings.Contains(chtoto, " + "):
		parts := strings.Split(chtoto, " + ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		num, err := strconv.Atoi(parts[0])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		result := strconv.Itoa(num) + parts[1]
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(chtoto, " - "):
		parts := strings.Split(chtoto, " - ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		num, err := strconv.Atoi(parts[0])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		result := strings.Replace(strconv.Itoa(num), parts[1], "", 1)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(chtoto, " * "):
		parts := strings.Split(chtoto, " * ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		num, err := strconv.Atoi(parts[0])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		result := strings.Repeat(parts[1], num)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	case strings.Contains(chtoto, " / "):
		parts := strings.Split(chtoto, " / ")
		if len(parts) != 2 {
			panic("неверное выражение")
		}
		num, err := strconv.Atoi(parts[1])
		if err != nil || num < 1 || num > 10 {
			panic("неверное выражение")
		}
		result := parts[0][:len(parts[0])/num]
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Println(result)

	default:
		panic("неверное вражение")
	}
}
