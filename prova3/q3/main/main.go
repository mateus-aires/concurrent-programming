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
	"fmt"
	"os"
	"path/filepath"
	"bufio"
)

func main() {
	var path string
	fmt.Print("Absolute path: ")
	fmt.Scanf("%s", &path)
	fmt.Println(path)

	directory := make(chan string)
	join := make(chan int)

	go navigate(directory, path)
	go readFile(directory, join)

	<- join
}


func navigate(ch chan<- string, root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// ch <- info.Name()
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
		// fmt.Println("File path: ", file)

		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		bufferedReader := bufio.NewReader(f)

		data, err := bufferedReader.ReadByte()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("First byte: %c\n", data)

		f.Close()
	}

	join_ch <- 1
}