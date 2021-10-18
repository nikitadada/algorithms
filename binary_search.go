package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(100)
	arr = arr[:10]

	fmt.Println("Сгенерирован псевдослучайный массив чисел:")
	sort.Ints(arr)
	printArray(arr)

	fmt.Print("Введите число для поиска в сгенерированном псевдослучайном массиве: ")
	searchInt := readInt()

	intsMap := map[int]int{}
	for i, v := range arr {
		intsMap[v] = i
	}

	for {
		arr = searh(arr, searchInt)
		if len(arr) == 1 {
			if arr[0] == searchInt {
				fmt.Print("Индекс запрашиваемого числа: ")
				println(intsMap[arr[0]])
				break
			} else {
				fmt.Print("Число не найдено ")
				break
			}
		}
	}
}

func searh(arr []int, n int) []int {
	mid := len(arr) / 2

	if arr[mid] > n {
		return arr[:mid]
	} else {
		return arr[mid:]
	}
}

func readInt() int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
	println()
}
