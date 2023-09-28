package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const (
	host        = "0.beevik-ntp.pool.ntp.org"
	reserveHost = "0.europe.pool.ntp.org"
)

func GetTime(address string) (time.Time, error) {
	return ntp.Time(address)
}

func main() {
	stdErrLog := log.New(os.Stderr, "", 1)
	NTPTime, err := GetTime(host)
	if err != nil {
		stdErrLog.Fatalln(err) // пишет лог в stderr и вызывает os.Exit(1)
		//либо вместо кастомного логера можно так:
		//fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
	}
	fmt.Println(NTPTime)
}
