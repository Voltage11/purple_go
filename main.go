package main

import (
	"fmt"
	"math/rand"

)

func main() {
	countItems := 10
	
	chRandom := make(chan int, countItems)
	chSqrt := make(chan int, countItems)

	go generateRandomToChannel(chRandom, countItems)
	
	for i := 0; i < countItems; i++ {
		select {
			case random := <-chRandom:
				go doSqrt(random, chSqrt)
		}
	}
	for i := 0; i < countItems; i++ {
		select {
			case sqrt := <-chSqrt:
				fmt.Print(sqrt, " ")
		}
	}

}

func generateRandomToChannel(ch chan int, countItems int) {
	for i := 0; i < countItems; i++ {
		ch <- rand.Intn(101)
	}
}

func doSqrt(num int, chSqrt chan int) {
	chSqrt <- num * num
}