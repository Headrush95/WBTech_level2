package dev02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
*/
var (
	invalidStringError      = errors.New("invalid input string")
	escapeRune         rune = 92
)

func UnpackString(src string) (string, error) {
	if src == "" || unicode.IsDigit([]rune(src)[0]) {
		return "", invalidStringError
	}
	var res strings.Builder
	input := []rune(src)
	var prevRune rune
	for idx, sym := range input {
		if unicode.IsDigit(sym) && (input[idx-1] != escapeRune || prevRune == escapeRune) {
			// если у нас две цифры подряд, то принимаем это за ошибку
			if unicode.IsDigit(input[idx-1]) && prevRune != input[idx-1] {
				return "", invalidStringError
			}

			// ошибку можно опустить, так как в IsDigit() уже проверили, что это цифра
			num, _ := strconv.Atoi(string(sym))
			for i := 0; i < num-1; i++ {
				res.WriteRune(prevRune)
			}
			continue
		}
		if sym == escapeRune {
			if idx != 0 && input[idx-1] == escapeRune {
				res.WriteRune(sym)
				prevRune = sym
				continue
			}
			continue
		}
		res.WriteRune(sym)
		prevRune = sym
	}

	return res.String(), nil
}
