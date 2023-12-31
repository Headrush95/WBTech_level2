Что выведет программа? Объяснить вывод программы. Рассказать про
внутреннее устройство слайсов и что происходит при передаче их в качестве
аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}
func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```
Ответ:
```
Вывод:
[3 2 3]
```
Чтобы разобраться с выводом программы лучше сначала взглянуть под капот структуры слайса:
```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
Здесь поле array - ссылка на массив с данными, len - текущая длина массива (кол-во элементов в нем), cap - вместимость
слайса, то есть граница, до которой он может расти без переалоцирования данных.

Возвращаясь к нашему примеру:
Если сразу указывать элементы при создании слайса, то поля len и cap будут равны (в нашем случае 3-м).
В функцию modifySlice мы передаем слайс по значению, то есть копируем его, но поскольку поле array является ссылкой, то
формально, подмассив передается по ссылке. Таким образом, при изменении элемента с индексом 0 мы его также меняем в
исходном слайсе. Однако, дальше происходит операция присоединения элемента append, которая инкрементирует поле len и
проверяет не стало ли оно больше значения поля cap. Если стало, то происходит переалокация данных и размер подмассива
увеличивается вдвое (после достижения 1024 элементов рост составляет 25%). В нашем случае именно это и происходит: мы
копируем слайс "s" с длиной и вместимостью равными 3-м, и при его росте, поскольку len=cap, создается новый подмассив с
длиной 4 и вместимостью 6, но так как мы имеем дело с копией слайса, то в исходном "s" это изменение не отражается.