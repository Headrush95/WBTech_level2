Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```
Ответ:
```
Вывод:
error
```
Возвращаясь в листингу 3 мы видим, что в функции test мы присваиваем полю tab значение типа customError.
При проверке err == nil, поскольку одно из полей не пустое, получим false. Чтобы это исправить надо из функции test()
возвращать интерфейсный тип error. А поскольку наш тип customError удовлетворяет ему (есть метод Error() string), то мы
без проблем можем возвращать наш тип ошибок из функции. 
