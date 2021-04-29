/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Race Condition — это недостаток, возникающий,
		когда время или порядок событий влияют на правильность программы.
	*/

	// for i := 0; i < 5; i++ {
	go func() {
		fmt.Printf("A->")
	}()

	go func() {
		fmt.Printf("B")
	}()
	// }
	time.Sleep(1 * time.Second)

	/*

		Output:
		BA->
		A->B
		A->B
		A->B
		BA->

		Поскольку результат не определен,
		он зависит от порядка выполнения горутин.

		Также Race Condition - может возникать, из-за того,
		что две Goroutine'ы "стучаться" в один общий ресурс:
		Отсутствует синхронизация - Mutex
	*/

	x := 0
	for i := 0; i < 10; i++ { // Ограничим в 10 итераций, иначе ПК отправится спать.
		go func() {
			x++
		}()

		go func() {
			if x%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
		}()
	}
}
