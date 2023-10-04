package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	input         = os.Stdin
	output        = os.Stdout
	fieldsFlag    = flag.String("f", "", "select fields (columns)")
	delimiterFlag = flag.String("d", "\t", "use a another delimiter")
	separatedFlag = flag.Bool("s", false, "only delimited lines")
)

type fieldRange struct {
	start int
	end   int
}

// parseFields парсит вводимы пользователем через флаг -f диапазоны столбцов для вывода
func parseFields(src string) ([]fieldRange, error) {
	fieldsArr := make([]fieldRange, 0, len(strings.Split(src, ",")))
	var inputField fieldRange
	var err error

	for _, field := range strings.Split(src, ",") {
		inputField = fieldRange{}

		if strings.Contains(field, "-") {
			tmp := strings.Split(field, "-")
			if len(tmp) != 2 {
				return nil, errors.New("invalid field range")
			}

			inputField.start, err = strconv.Atoi(tmp[0])
			if err != nil {
				return nil, err
			}
			inputField.end, err = strconv.Atoi(tmp[1])
			if err != nil {
				return nil, err
			}

			fieldsArr = append(fieldsArr, inputField)
			continue
		}

		fieldNumb, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}

		inputField.start, inputField.end = fieldNumb, fieldNumb
		fieldsArr = append(fieldsArr, inputField)
	}

	return fieldsArr, nil
}

func main() {
	flag.Parse()

	// если нет указанных полей, то просим указать и завершаем работу
	if *fieldsFlag == "" {
		fmt.Println("Please enter fields numbers")
		os.Exit(1)
	}

	// парсим номера столбцов из флага -f
	fieldsNumb, err := parseFields(*fieldsFlag)
	if err != nil {
		log.Fatalln(err)
	}

	var countOfLines int
	var line string

	// считываем строки из os.Stdin
	fmt.Print("Enter count of lines: ")
	_, err = fmt.Scanln(&countOfLines)
	if err != nil {
		log.Fatalln(err)
	}

	inputArray := make([][]string, 0, countOfLines)

	fmt.Println()
	fmt.Println("Enter lines:")
	rdr := bufio.NewReader(input)

	for i := 0; i < countOfLines; i++ {
		line, _ = rdr.ReadString('\n')
		line = strings.TrimRight(line, "\n\r")
		// inputArray - массив разбитых разделителем строк
		inputArray = append(inputArray, strings.Split(line, *delimiterFlag))
	}
	// завершаем чтение

	wr := bufio.NewWriter(output)
	defer wr.Flush()
	newLine := strings.Builder{}

	// итерируемся по входных строкам и пытаемся вывести запрошенные столбцы
	for _, lineArr := range inputArray {
		if *separatedFlag && len(lineArr) == 1 {
			continue
		}

		for i, columnRange := range fieldsNumb {
			for curColumn := columnRange.start - 1; curColumn < columnRange.end; curColumn++ {
				if curColumn > len(lineArr)-1 {
					newLine.WriteString("---")
					// у последнего столбца разделитель не пишем
					if i == len(fieldsNumb)-1 {
						continue
					}

					newLine.WriteString(*delimiterFlag)
					continue
				}

				newLine.WriteString(lineArr[curColumn])
				// у последнего столбца разделитель не пишем
				if i == len(fieldsNumb)-1 {
					continue
				}
				newLine.WriteString(*delimiterFlag)
			}
		}
		newLine.WriteRune('\n')

		_, err = wr.WriteString(newLine.String())
		if err != nil {
			log.Fatalln(err)
		}
		newLine.Reset()
	}

}
