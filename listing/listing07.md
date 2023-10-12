Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```
Ответ:
```
Вывод: будут выведены в случайном порядке числа, переданные в "a" и "b", затем будут бесконечно выводиться нули. Каждое
число на новой строчке.
```
В данном случае, функция asChan создает горутину-чтеца и возвращает канал, в который горутина пишет. После завершения
чтения канал закрывается (!). Эта функция вызывается дважды - для создания переменных "a" и "b". Дальше вызывается функция
соединения каналов в один. И в конце происходит чтение из объединенного канала.

В функции merge есть недостатки: во-первых, канал "c" не закрывается, то есть в вызывающей горутине это может вызвать
блокировку. Во-вторых, в блоке select не происходит проверка на закрытие канала, а значит, даже если канал закрыт, будет
приходить значение по умолчанию (свойство каналов) - в нашем случае 0. Что и происходит.