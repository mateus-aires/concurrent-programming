package main

/*
two-phase sleep
Crie um programa que recebe um número inteiro n como argumento e cria n goroutines.
Cada uma dessas goroutines deve dormir por um tempo aleatório de no máximo 5 segundos.
Depois que acordar, cada thread deve sortear um outro número aleatório s (entre 0 e 10).
Somente depois de todas as n goroutine terminarem suas escolhas
(ou seja, ao fim da primeira fase), começamos a segunda fase.
Nesta segunda fase, a n-ésima goroutine criada deve dormir pelo tempo s escolhido
pela goroutine n - 1 (faça a contagem de maneira modular, ou seja, a primeira goroutine
dorme conforme o número sorteado pela última).
*/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
func main() {
	var goroutines_to_create int
	fmt.Print("n: ")
	fmt.Scanf("%d", &goroutines_to_create)

	var first_phase_wg sync.WaitGroup
	var second_phase_wg sync.WaitGroup

	first_phase_wg.Add(goroutines_to_create)
	second_phase_wg.Add(goroutines_to_create)

	var channels = make([]chan int, goroutines_to_create)
	for i := 0; i < goroutines_to_create; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < goroutines_to_create; i++ {
		var previous_index int
		var previous chan int
		next := channels[i]
		if i == 0 {
			previous_index = len(channels) - 1
		} else {
			previous_index = i - 1
		}

		fmt.Printf("i: %d, previous: %d\n", i, previous_index)
		previous = channels[previous_index]
		go func(thread_id int) {
			two_phase_sleep(thread_id, previous, next, &first_phase_wg, &second_phase_wg)
		}(i)
	}

	second_phase_wg.Wait()
	fmt.Printf("%d\n", goroutines_to_create)
}

func two_phase_sleep(thread_id int, previous chan int, next chan int, first_wg *sync.WaitGroup, second_wg *sync.WaitGroup) {
	first_phase(thread_id, next, first_wg)
	first_wg.Wait()

	second_phase(thread_id, previous, second_wg)

}

func first_phase(thread_id int, next chan int, first_wg *sync.WaitGroup) {
	defer first_wg.Done()

	rand.Seed(time.Now().UnixNano())
	sleep_time := rand.Intn(5)

	fmt.Printf("[#%d] First sleep %d seconds\n", thread_id, sleep_time)
	time.Sleep(time.Duration(sleep_time) * time.Second)
	fmt.Printf("[#%d] First sleep finished\n", thread_id)

	n := rand.Intn(10)
	fmt.Printf("[#%d] Number drawn: %d\n", thread_id, n)
	next <- n
}

func second_phase(thread_id int, previous chan int, second_wg *sync.WaitGroup) {
	defer second_wg.Done()

	sleep_time := <-previous

	fmt.Printf("[#%d] Second sleep %d seconds\n", thread_id, sleep_time)
	time.Sleep(time.Duration(sleep_time) * time.Second)
	fmt.Printf("[#%d] Second sleep finished\n", thread_id)
}
