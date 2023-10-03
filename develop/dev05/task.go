package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
)

var (
	afterFlag      = flag.Int("A", -1, "print +N lines after match")
	beforeFlag     = flag.Int("B", -1, "print +N lines before match")
	contextFlag    = flag.Int("C", -1, "print ±N lines around match")
	countFlag      = flag.Bool("c", false, "count of lines")
	ignoreCaseFlag = flag.Bool("i", false, "ignore case")
	invertFlag     = flag.Bool("v", false, "instead of matching, exclude")
	fixedFlag      = flag.Bool("F", false, "exact match to string, not a pattern")
	lineNumbFlag   = flag.Bool("n", false, "print line number")
)

func readInput(input string) (map[int]string, error) {
	f, err := os.Open(input)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	scn := bufio.NewScanner(f)

	// считываем построчно в мапу, где ключ - номер строки, а значение - строка
	inputLines := make(map[int]string, 100)
	lineIdx := 1
	for scn.Scan() {
		inputLines[lineIdx] = scn.Text()
		lineIdx++
	}
	if err = scn.Err(); err != nil {
		return nil, err
	}
	return inputLines, nil
}

func writeOutput(output string, src map[int]string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	wr := bufio.NewWriter(f)
	defer wr.Flush()

	found, err := findEntry(src)
	if *countFlag {
		_, err = wr.WriteString(fmt.Sprintf("Total count of lines containing the pattern in %s - %d\n", inputFile, len(found)))
		if err != nil {
			return err
		}
		return nil
	}

	var printedLine strings.Builder
	// пишем в файл
	for _, entry := range found {
		for lineNum, line := range entry {
			if *lineNumbFlag {
				printedLine.WriteString(strconv.Itoa(lineNum))
				printedLine.WriteString(" ")
			}
			printedLine.WriteString(line)
			_, err = wr.WriteString(printedLine.String() + "\n")
			if err != nil {
				return err
			}
			printedLine.Reset()
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func findEntry(src map[int]string) ([]map[int]string, error) {
	result := make([]map[int]string, 0, 10)
	pattern := flag.Arg(0)
	if *ignoreCaseFlag && !*fixedFlag {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}

	switch {
	// ищем строки без использования регулярок
	case *fixedFlag:
		for lineNum, line := range src {
			if strings.Contains(line, pattern) {
				result = append(result, getAdjacentLines(src, lineNum))
			}
		}

	// с регулярками
	default:
		regExp, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		for lineNum, line := range src {
			if (regExp.MatchString(line) && !*invertFlag) || (!regExp.MatchString(line) && *invertFlag) {
				result = append(result, getAdjacentLines(src, lineNum))
			}
		}
	}

	return result, nil
}

func getAdjacentLines(src map[int]string, lineNumb int) map[int]string {
	lower := lineNumb
	upper := lineNumb

	// TODO убрать повторения кода
	if *contextFlag > 0 {
		if lineNumb-*contextFlag < 1 {
			lower = 1
		} else {
			lower = lineNumb - *contextFlag
		}

		if lineNumb+*contextFlag > len(src) {
			upper = len(src)
		} else {
			upper = lineNumb + *contextFlag
		}
	} else if *afterFlag > 0 {
		if lineNumb+*afterFlag > len(src) {
			upper = len(src)
		} else {
			upper = lineNumb + *afterFlag
		}
	} else if *beforeFlag > 0 {
		if lineNumb-*beforeFlag < 0 {
			lower = 1
		} else {
			lower = lineNumb - *beforeFlag
		}
	}

	adjacentLines := make(map[int]string, upper-lower+1)
	for i := lower; i <= upper; i++ {
		adjacentLines[i] = src[i]
	}

	return adjacentLines
}

func main() {
	flag.Parse()
	src, err := readInput(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeOutput(outputFile, src)
	if err != nil {
		log.Fatalln(err)
	}
}
