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
	var goroutinesToCreate int
	fmt.Print("n: ")
	fmt.Scanf("%d", &goroutinesToCreate)

	//makes main goroutine wait until all created goroutines finished
	var secondPhaseWG sync.WaitGroup
	secondPhaseWG.Add(goroutinesToCreate)

	var channels = make([]chan int, goroutinesToCreate)
	for i := 0; i < goroutinesToCreate; i++ {
		channels[i] = make(chan int)
	}

	for i := 0; i < goroutinesToCreate; i++ {
		go func(goRoutineID int) {
			twoPhaseSleep(goRoutineID, channels[goRoutineID], &secondPhaseWG)
		}(i)
	}

	// Get values for the second phase sleep
	var secondSleepTimes = make([]int, goroutinesToCreate)
	for i := 0; i < goroutinesToCreate; i++ {
		secondSleepTimes[i] = <-channels[i]
	}
	fmt.Printf("[Main goroutine] Second phase sleep times: %v\n", secondSleepTimes)

	// Send value drawn by the n - 1 goroutine
	for i := 0; i < goroutinesToCreate; i++ {
		sleepTimeIndex := getSecondPhaseSleepTimeIndex(i, goroutinesToCreate)
		channels[i] <- secondSleepTimes[sleepTimeIndex]
	}

	secondPhaseWG.Wait()
	fmt.Printf("[Main goroutine] n: %d\n", goroutinesToCreate)
}

func twoPhaseSleep(id int, ch chan int, secondWG *sync.WaitGroup) {
	firstPhase(id, ch)
	secondPhase(id, ch, secondWG)
}

func firstPhase(id int, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	sleepTime := rand.Intn(5)

	fmt.Printf("[#%d] First sleep %d seconds\n", id, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("[#%d] First sleep finished\n", id)

	ch <- rand.Intn(10)
}

func secondPhase(id int, ch chan int, second_wg *sync.WaitGroup) {
	defer second_wg.Done()
	sleepTime := <-ch
	fmt.Printf("[#%d] Second sleep %d seconds\n", id, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("[#%d] Second sleep finished\n", id)
}

func getSecondPhaseSleepTimeIndex(goroutineID int, goroutinesCreated int) int {
	var index int
	if goroutineID == 0 {
		index = goroutinesCreated - 1
	} else {
		index = goroutineID - 1
	}
	return index
}
