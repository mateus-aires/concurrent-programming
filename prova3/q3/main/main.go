package main

/*
pipeline
Crie um programa organizado como um pipeline de goroutines. Esse programa deve
receber como argumento um caminho absoluto para um diretório. Uma goroutine deve
navegar na árvore que tem como raiz o diretório passado como argumento. Essa
goroutine deve passar para uma próxima goroutine do pipeline o nome dos arquivos
encontrados na busca dos diretórios, ou seja, ignore os diretórios. Esta segunda
goroutine deve ler o primeiro byte de conteúdo de cada um desses arquivos e
escrever na saída padrão o nome dos arquivos que tem esse valor do primeiro byte
sendo par.
*/

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Print("Absolute path: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	path := scanner.Text()

	directory := make(chan string)
	join := make(chan int)

	go navigate(directory, path)
	go readFile(directory, join)

	<-join
}

func navigate(ch chan<- string, root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ch <- path
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	close(ch)
}

func readFile(ch <-chan string, join_ch chan<- int) {
	for file := range ch {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		bufferedReader := bufio.NewReader(f)

		data, err := bufferedReader.ReadByte()
		if err != nil {
			fmt.Println(err)
		}

		if data%2 == 0 {
			fmt.Printf("File: %s | First byte: %c | First Byte Value: %v\n", file, data, data)
		}

		f.Close()
	}

	join_ch <- 1
}
