package pattern

import (
	"fmt"
	"log"
)

/*
Реализовать паттерн "Цепочка вызовов", объяснить применимость паттерна, плюсы и минусы, а
также реальные примеры его использования на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Цепочка вызовов - это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по
цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос
дальше по цепи.

Применимость:
1) когда программа содержит несколько объектов, способных обработать тот или иной запрос, однако заранее неизвестно какой
запрос придет и какой обработчик понадобится;
2) когда важно, чтобы обработчики выполнялись один за другим в строгом порядке (последовательная обработка чего-либо);
3) когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы:
- уменьшает зависимость между клиентом и обработчиком;
- реализует принцип единственной ответственности, так как мы реализуем последовательную обработку, где каждый обработчик
делает что-то одно;
- реализует принцип открытости/закрытости, так как новое поведение будет добавлено через новый обработчик.

Минусы:
- запрос может остаться никем не обработанным, если нет подходящего обработчика.

Примеры:
1) Допустим, у нас есть веб-сервис, который позволяет хранить мультимедийные файлы. Можно организовать следующую цепочку
обработчиков: проверка типа файла checkType(), чтобы убедиться, что файл имеет валидный тип => выбор соответствующего
алгоритма сжатия файла selectCompressionType(), так как для разных файлов он будет разный => загрузка в соответствующее
типу хранилище upload(), чтобы файлы у нас были лучше структурированы.

Реализация:
Рассмотрим банк. Допустим, у банка есть условие, что при покупке, если средств на основном счете не хватает, то он спишет
 их с любого другого, где средств достаточно.
*/

// PaymentAccount - интерфейс нашего платежного аккаунта. Если на нем недостаточно средств для покупки, то он попробует
// списать средства со следующего аккаунта
type PaymentAccount interface {
	SetNextPaymentAccount(next PaymentAccount)
	Pay(amount int)
}

type BankAccount struct {
	balance int
	next    PaymentAccount
}

func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{balance: balance, next: nil}
}

func (b *BankAccount) SetNextPaymentAccount(next PaymentAccount) {
	b.next = next
}

func (b *BankAccount) Pay(amount int) {
	if b.balance < amount {
		// если средств недостаточно и есть другие аккаунты, то пробуем оплатить используя их
		if b.next != nil {
			b.next.Pay(amount)
			return
		}
		log.Println("not enough money")
		return
	}
	b.balance -= amount
}

func (b *BankAccount) GetBalance() int {
	return b.balance
}

func example() {
	mainAccount := NewBankAccount(100)
	secondaryAccount := NewBankAccount(500)

	mainAccount.SetNextPaymentAccount(secondaryAccount)
	mainAccount.Pay(200)
	fmt.Println(mainAccount.GetBalance())
	fmt.Println(secondaryAccount.GetBalance())
}