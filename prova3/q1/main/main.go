package main

/*
Fork-sleep-join
Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos.
A main-goroutine deve esperar todas as goroutines filhas terminarem de executar para
em seguida escrever na saída padrão o valor de n.
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var goroutines_to_create int

	fmt.Print("n: ")
	fmt.Scanf("%d", &goroutines_to_create)
	ch := make(chan bool)

	for i := 0; i < goroutines_to_create; i++ {
		go random_sleep(i, ch)
	}

	for i := 0; i < goroutines_to_create; i++ {
		<-ch
	}
	fmt.Printf("%d\n", goroutines_to_create)
}

func random_sleep(thread_id int, done chan bool) {
	rand.Seed(time.Now().UnixNano())
	sleep_time := rand.Intn(5)
	fmt.Printf("[#%d] Going to sleep for seconds %d\n", thread_id, sleep_time)
	time.Sleep(time.Duration(sleep_time) * time.Second)
	fmt.Printf("[#%d] Sleep finished\n", thread_id)
	done <- true
}
