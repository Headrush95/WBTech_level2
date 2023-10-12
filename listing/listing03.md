Что выведет программа? Объяснить вывод программы. Объяснить внутреннее
устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

```

Ответ:
```
Вывод:
<nil>
false
```
Чтобы лучше понять почему так, надо заглянуть под капот интерфейса. Взглянем на структуру iface из пакета runtime:

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

Здесь tab — это указатель на Interface Table или itable — структуру, которая хранит некоторые метаданные о типе и
список методов, используемых для удовлетворения интерфейса. data — указывает на фактическую переменную с конкретным
(статическим) типом.
В примере выше мы помещаем в поле tab значение *os.PathError, а в data nil. При сравнении структуры err с nil, если
хотя бы одно из полей имеет не нулевое значение, результат будет false.

Пустой интерфейс не имеет поля tab и хранить только данные, но не тип с список методов. Поэтому пустому интерфейсу
удовлетворяет любой тип.