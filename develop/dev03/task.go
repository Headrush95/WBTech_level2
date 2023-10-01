package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

TODO
Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

*/

const (
	inputFile = "input.txt"
	outFile   = "output.txt"
)

var (
	columnNumbForSortingFlag = flag.Int("k", -1, "specifying the column to sort")
	numericValueSortingFlag  = flag.Bool("n", false, "sort by numeric value")
	reverseSortFlag          = flag.Bool("r", false, "sort in reverse order")
	notShowDuplicates        = flag.Bool("u", false, "do not print duplicate lines")
)

// readInput считывает данные из указанного файла
func readInput(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	res := make([]string, 0, 100)

	for scn.Scan() {
		res = append(res, scn.Text())
	}
	if err = scn.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// baseSort выполняет обычную сортировку
func baseSort(input []string) {
	slices.Sort(input)
}

// reverse переворачивает слайс
func reverse(input []string) {
	length := len(input) - 1
	for i := 0; i <= length/2; i++ {
		input[i], input[length-i] = input[length-i], input[i]
	}
}

// writeOutput печатает результат сортировки в указанный файл
func writeOutput(output []string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	wr := bufio.NewWriter(f)
	defer wr.Flush()
	for _, line := range output {
		_, err = wr.WriteString(line + "\n")
		if err != nil {
			return nil
		}
	}
	return nil
}

// deleteDuplicates удаляет дубликаты, если активен флаг notShowDuplicates
func deleteDuplicates(input []string) []string {
	if len(input) == 0 {
		return input
	}

	resSet := make(map[string]struct{}, len(input))
	result := make([]string, 0, len(input))

	for _, line := range input {
		if _, ok := resSet[line]; !ok {
			resSet[line] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}

// numericSort пытается привести строки к числу и отсортировать
func numericSort(input []string) {
	sort.Slice(input, func(i, j int) bool {
		iNum, err := strconv.Atoi(input[i])
		if err != nil {
			return input[i] < input[j]
		}

		jNum, err := strconv.Atoi(input[j])
		if err != nil {
			return input[i] < input[j]
		}
		return iNum < jNum
	})
}

// sortByColumn сортирует по указанной колонке. Если в строке колонка пустая, то строка оказывается внизу списка.
func sortByColumn(input []string, columnIdx int, needSortNumeric bool) {
	sort.Slice(input, func(i, j int) bool {
		if len(strings.Fields(input[i])) < columnIdx+1 {
			return false
		}
		if len(strings.Fields(input[j])) < columnIdx+1 {
			return true
		}

		if needSortNumeric {
			iNum, iErr := strconv.Atoi(strings.Fields(input[i])[columnIdx])

			jNum, jErr := strconv.Atoi(strings.Fields(input[j])[columnIdx])

			if iErr != nil && jErr != nil {
				return strings.Fields(input[i])[columnIdx] < strings.Fields(input[j])[columnIdx]
			}
			if iErr != nil {
				return false
			}
			if jErr != nil {
				return true
			}

			return iNum < jNum
		}

		return strings.Fields(input[i])[columnIdx] < strings.Fields(input[j])[columnIdx]
	})
}

func main() {
	flag.Parse()
	inputData, err := readInput(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	if *notShowDuplicates {
		inputData = deleteDuplicates(inputData)
	}

	if *columnNumbForSortingFlag >= 0 {
		sortByColumn(inputData, *columnNumbForSortingFlag, *numericValueSortingFlag)
	} else if *numericValueSortingFlag {
		numericSort(inputData)
	} else {
		baseSort(inputData)
	}

	if *reverseSortFlag {
		reverse(inputData)
	}
	err = writeOutput(inputData, outFile)
	if err != nil {
		log.Fatalln(err)
	}
}
