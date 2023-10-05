package main

import (
	"bufio"
	"errors"
	"github.com/google/gops/goprocess"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут
быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качестве
аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в
формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах
(linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*Шелл — это обычная консольная программа, которая будучи
запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный
сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).

*/

// changeDirectoryCommand - меняет рабочую папку на указанную
func changeDirectoryCommand(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}

// presentWorkingDirectoryCommand - выводит текущую рабочую папку
func presentWorkingDirectoryCommand(wr bufio.Writer) error {
	defer wr.Flush()
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = wr.WriteString(currentDir + "\n")
	return err
}

// echoCommand - выводит переданный текст на стандартное устройство вывода
func echoCommand(wr bufio.Writer, src []string) error {
	defer wr.Flush()
	var err error

	for _, val := range src {
		_, err = wr.WriteString(val + " ")
		if err != nil {
			return err
		}
	}
	_, err = wr.WriteString("\n")

	return err
}

// killProcessCommand - завершает процесс с указанным PID
func killProcessCommand(id string) error {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	process, err := os.FindProcess(uid)
	if err != nil {
		return err
	}

	err = process.Kill()
	return err
}

// processStatusCommand - выводит все запущенные процессы на хосте
func processStatusCommand(wr bufio.Writer) error {
	processes, err := ps.Processes()
	if err != nil {
		return err
	}
	defer wr.Flush()

	processInfo := strings.Builder{}
	for _, process := range processes {
		processInfo.WriteString("Exec: ")
		processInfo.WriteString(process.Executable())

		processInfo.WriteString(" PID: ")
		processInfo.WriteString(strconv.Itoa(process.Pid()))

		processInfo.WriteString(" PPID: ")
		processInfo.WriteString(strconv.Itoa(process.PPid()))

		processInfo.WriteString("\n")
		_, err = wr.WriteString(processInfo.String())
		processInfo.Reset()

		if err != nil {
			return err
		}
	}

	return nil
}

// goProcessStatusCommand - выводит все запущенные go процессы на хосте
func goProcessStatusCommand(wr bufio.Writer) error {
	defer wr.Flush()

	processes := goprocess.FindAll()
	processInfo := strings.Builder{}
	var err error
	for _, process := range processes {
		processInfo.WriteString("Exec: ")
		processInfo.WriteString(process.Exec)

		processInfo.WriteString(" Path: ")
		processInfo.WriteString(process.Path)

		processInfo.WriteString(" Version: ")
		processInfo.WriteString(process.BuildVersion)

		processInfo.WriteString(" PID: ")
		processInfo.WriteString(strconv.Itoa(process.PID))

		processInfo.WriteString(" PPID: ")
		processInfo.WriteString(strconv.Itoa(process.PPID))

		processInfo.WriteString("\n")
		_, err = wr.WriteString(processInfo.String())
		processInfo.Reset()
		if err != nil {
			return err
		}
	}

	return nil
}

// parseCommand - вспомогательная функция для извлечения параметров команды
func parseCommand(input string) []string {
	args := strings.Fields(input)
	return args
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var command string
	var err error
	var commandArgs []string

	for {
		wr.WriteString(">>> ")
		err = wr.Flush()
		if err != nil {
			log.Println(err)
			continue
		}

		command, err = rdr.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		commandArgs = parseCommand(strings.Trim(command, "\n\r"))

		switch commandArgs[0] {
		case "cd":
			if len(commandArgs) != 2 {
				err = errors.New("invalid count of arguments")
				break
			}

			err = changeDirectoryCommand(commandArgs[1])
		case "pwd":
			if len(commandArgs) != 1 {
				err = errors.New("invalid count of arguments")
				break
			}

			err = presentWorkingDirectoryCommand(*wr)
		case "echo":
			err = echoCommand(*wr, commandArgs[1:])
		case "kill":
			if len(commandArgs) != 2 {
				err = errors.New("invalid count of arguments")
				break
			}

			err = killProcessCommand(commandArgs[1])
		case "gops":
			if len(commandArgs) != 1 {
				err = errors.New("invalid count of arguments")
				break
			}

			err = goProcessStatusCommand(*wr)
		case "ps":
			if len(commandArgs) != 1 {
				err = errors.New("invalid count of arguments")
				break
			}

			err = processStatusCommand(*wr)
		case "exit":
			return

		default:
			log.Println("unknown command")
		}

		if err != nil {
			log.Println(err)
		}
	}
}
