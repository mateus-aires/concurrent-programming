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
	var goRoutinesToCreate int
	fmt.Print("n: ")
	fmt.Scanf("%d", &goRoutinesToCreate)

	var firstPhaseWg sync.WaitGroup
	var secondPhaseWg sync.WaitGroup
	firstPhaseWg.Add(goRoutinesToCreate)
	secondPhaseWg.Add(goRoutinesToCreate)

	var channels = make([]chan int, goRoutinesToCreate)

	for i := 0; i < goRoutinesToCreate; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < goRoutinesToCreate; i++ {
		nextIndex := getNextIndex(i, goRoutinesToCreate)
		nextCh := channels[nextIndex]
		go func(thread_id int) {
			first_phase(thread_id, nextCh, &firstPhaseWg, goRoutinesToCreate)
		}(i)
	}

	var secondSleepTimes = make([]int, goRoutinesToCreate)

	for i := 0; i < goRoutinesToCreate; i++ {
		ch := channels[i]
		secondSleepTimes[i] = <-ch
		close(ch)
	}

	firstPhaseWg.Wait()

	fmt.Print("End of first phase.\n")

	for i := 0; i < goRoutinesToCreate; i++ {
		go func(threadId int) {
			second_phase(threadId, secondSleepTimes[threadId], &secondPhaseWg)
		}(i)
	}
	secondPhaseWg.Wait()

	// second_phase_wg.Wait()
	fmt.Print("Finished.")
}

func first_phase(threadId int, next chan int, firstWg *sync.WaitGroup, totalRoutines int) {
	defer firstWg.Done()

	rand.Seed(time.Now().UnixNano())
	sleepTime := rand.Intn(5)

	fmt.Printf("[#%d] First sleep %d seconds\n", threadId, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("[#%d] First sleep finished\n", threadId)

	n := rand.Intn(10)
	nextIndex := getNextIndex(threadId, totalRoutines)
	fmt.Printf("[#%d] Routine #%d will sleep for %d seconds.\n", threadId, nextIndex, n)
	next <- n
}

func second_phase(thread_id int, sleepTime int, second_wg *sync.WaitGroup) {
	defer second_wg.Done()

	fmt.Printf("[#%d] Second sleep %d seconds\n", thread_id, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("[#%d] Second sleep finished\n", thread_id)
}

func getNextIndex(i int, total int) int {
	nextIndex := 0
	if i == (total - 1) {
		nextIndex = 0
	} else {
		nextIndex = i + 1
	}
	return nextIndex
}
