package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты
целиком.
*/
var (
	urlFlag            = flag.String("u", "", "url")
	outputFileNameFlag = flag.String("n", "output", "output file name")
)

func wget(url, outputFile string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	err = writeRespToFile(resp, outputFile)

	return err
}

func writeRespToFile(resp *http.Response, outputFile string) error {
	f, err := os.Create(outputFile + ".html")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	wr := bufio.NewWriter(f)
	defer func() {
		if err := wr.Flush(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = io.Copy(wr, resp.Body)
	return err
}

func main() {
	flag.Parse()
	if *urlFlag == "" {
		log.Fatalln("invalid url")
	}

	if err := wget(*urlFlag, *outputFileNameFlag); err != nil {
		log.Fatalln(err)
	}
}
