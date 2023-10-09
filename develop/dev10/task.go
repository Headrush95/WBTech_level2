package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeoutFlag = flag.String("timeout", "10s", "server connection timeout")
)

type Socket struct {
	host string
	port string
}

type Connection struct {
	socket  Socket
	timeout time.Duration
}

func parseTimeout() time.Duration {
	timeout, err := time.ParseDuration(*timeoutFlag)
	if err != nil {
		log.Fatalln(err)
	}

	return timeout
}

func parseConnectionParams() Connection {
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalln(errors.New("[error] wrong number of arguments, need 2"))
	}

	timeout := parseTimeout()

	port := args[1]
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalln(fmt.Errorf("invalid port %s", port))
	}

	return Connection{
		socket: Socket{
			host: args[0],
			port: port,
		},
		timeout: timeout,
	}
}

func readUserInput(conn *net.Conn) {

	var line string
	var err error

	rdr := bufio.NewReader(os.Stdin)

	for {
		line, err = rdr.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		_, err = (*conn).Write([]byte(line))
		if err != nil {
			(*conn).Close()
			log.Fatalln(err)
		}
	}

}

func main() {
	flag.Parse()
	timeout := parseTimeout()
	connParams := parseConnectionParams()

	serverSocket := net.JoinHostPort(connParams.socket.host, connParams.socket.port)
	serverConn, err := net.DialTimeout("tcp", serverSocket, timeout)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintf(os.Stdout, "Connected to %s\n", serverSocket)
	if err != nil {
		log.Fatalln(err)
	}
	go readUserInput(&serverConn)

	// завершение программы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)

	<-quit
}
