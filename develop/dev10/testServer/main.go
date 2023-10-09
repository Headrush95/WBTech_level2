package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

/*
Вспомогательный сервис для подключения клиентов
*/

var (
	host = "localhost"
	port = "3000"
	out  = os.Stdout
)

func handleClient(client *net.Conn) {
	clientAddr := (*client).RemoteAddr().String()

	defer func() {
		err := (*client).Close()
		if err != nil {
			log.Println(err)
		}
		_, err = fmt.Fprintf(out, "[info] %s was disconnected...", clientAddr)
		if err != nil {
			log.Println(err)
		}
	}()

	_, err := fmt.Fprintf(out, "[info] %s was connected...\n", clientAddr)
	if err != nil {
		log.Println(err)
		return
	}

	// создаем читателя для получения сообщений от конкретного пользователя
	scanner := bufio.NewScanner(*client)

	for scanner.Scan() {
		_, err = fmt.Fprintf(out, "[%s] %s\n", clientAddr, scanner.Text())
		if err != nil {
			log.Println(err)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Println(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = fmt.Fprintf(out, "[info] Telnet server started...\n")

	for {
		// ждем подключения нового клиента
		client, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleClient(&client)
	}

}
