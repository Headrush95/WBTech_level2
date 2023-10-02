package dev04

import (
	"fmt"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortSliceOfRunes(word string) string {

	res := []rune(word)
	slices.Sort(res)
	return string(res)
}

// Поскольку в итоговой мапе у нас по условиям множества - удаляем повторяющиеся слова
func deleteDuplicates(src []string) []string {
	resultMap := make(map[string]struct{}, len(src))
	for _, word := range src {
		resultMap[word] = struct{}{}
	}
	result := make([]string, 0, len(resultMap))
	for k := range resultMap {
		result = append(result, k)
	}
	slices.Sort(result)
	return result
}

func findAnagram(input []string) map[string][]string {
	// длинна len(input) для того, чтобы поменять ключи с отсортированных рун на первые встретившиеся слова
	result := make(map[string][]string, len(input))

	for _, word := range input {
		word = strings.ToLower(word)
		sortedWord := sortSliceOfRunes(word)

		if _, ok := result[sortedWord]; !ok {
			result[sortedWord] = make([]string, 0, 10)
			result[sortedWord] = append(result[sortedWord], word)
			continue
		}
		result[sortedWord] = append(result[sortedWord], word)
	}
	resultLength := len(result)
	currentMapPosition := 0
	for sortedRunes, sliceOfWords := range result {
		if currentMapPosition == resultLength {
			break
		}

		// если у нас множество из одного элемента, то удаляем его
		if len(sliceOfWords) == 1 {
			delete(result, sortedRunes)
		}
		// меняем ключи с сортированных рун на первые встретившиеся слова
		result[result[sortedRunes][0]] = deleteDuplicates(result[sortedRunes])
		delete(result, sortedRunes)
		currentMapPosition++
	}
	return result
}

func example() {
	input := []string{"тяпка", "пятак", "листок", "пятка", "слиток", "тяпка", "столик"}
	fmt.Println(findAnagram(input))
}
