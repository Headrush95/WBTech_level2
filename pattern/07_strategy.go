package pattern

import "fmt"

/*
Реализовать паттерн "Стратегия", объяснить применимость паттерна, плюсы и минусы, а
также реальные примеры его использования на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия - это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
из них в собственный класс. После чего, алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применимость:
1) когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта;
2) когда есть множество похожих объектов, отличающихся только некоторым поведением;
3) когда нужно скрыть детали реализации алгоритмов для других объектов (классов);
4) когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора
представляет вариацию алгоритма.

Плюсы:
- горячая замена алгоритмов на лету;
- изолирует код и данные алгоритмов от остальных объектов (классов);
- уход от наследования к делегированию;
- реализует принцип открытости/закрытости.

Минусы:
- усложняет программу за счет дополнительных объектов (классов);
- клиент должен знать, в чем разница между стратегиями, чтобы выбрать подходящую.

Примеры:
1) Допустим, у нас есть приложение навигации. Сначала мы строили маршрут только для пользователей на машине, но потом
решили добавить еще прокладку маршрута на общественном транспорте и для пеших прогулок. С каждым новым вариантом код
основного объекта увеличивался вдвое. Его стало трудно читать и ориентироваться в нем. В данном случае, нам очень поможет
паттерн Стратегия: мы объединим наши алгоритмы прокладывания пути с помощью интерфейса Navigator и вынести их в отдельный
пакет. Таким образом, клиент не знает о внутренней реализации прокладывания маршрутов и может быстро переключаться между
разными алгоритмами.

Реализация:
Допустим, у нас есть сервис, который в зависимости от платформы имеет разные механизмы аутентификации (Basic, OAuth, Bearer).

*/

type User struct {
	platform   string
	authMethod AuthorizationStrategy
}

func (u *User) SetAuthMethod() {
	if u.platform == "mobile" {
		u.authMethod = &OAuth{}
		return
	} else if u.platform == "web" {
		u.authMethod = &BasicAuth{}
		return
	}
	u.authMethod = &BearerAuth{}
}

func (u *User) Authorize() {
	u.authMethod.Auth()
}

// AuthorizationStrategy - реализация стратегии для сервиса аутентификации
type AuthorizationStrategy interface {
	Auth()
}

type BasicAuth struct {
}

func (ba *BasicAuth) Auth() {
	fmt.Println("Authorized using Basic authorization")

}

type OAuth struct {
}

func (oa *OAuth) Auth() {
	fmt.Println("Authorized using OAuth")
}

type BearerAuth struct {
}

func (b *BearerAuth) Auth() {
	fmt.Println("Authorized using Bearer token")
}

func strategyExample() {
	user := User{platform: "mobile"}
	user.SetAuthMethod()
	user.Authorize()
}
