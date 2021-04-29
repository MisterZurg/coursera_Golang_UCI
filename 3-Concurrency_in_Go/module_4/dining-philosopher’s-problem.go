/*
Implement the dining philosopher’s problem
with the following constraints/modifications.

	1)	There should be 5 philosophers sharing chopsticks,
	with one chopstick between each adjacent pair of philosophers.

	2)	Each philosopher should eat only 3 times
	(not in an infinite loop as we did in lecture)

	3)	The philosophers pick up the chopsticks in any order,
	not lowest-numbered first (which we did in lecture).

	4)	In order to eat, a philosopher must get permission from a host
	which executes in its own goroutine.

	5)	The host allows no more than 2 philosophers to eat concurrently.

	6)	Each philosopher is numbered, 1 through 5.

	7)	When a philosopher starts eating (after it has obtained necessary locks)
	it prints “starting to eat <number>” on a line by itself,
	where <number> is the number of the philosopher.

	8)	When a philosopher finishes eating (before it has released its locks)
	it prints “finishing eating <number>” on a line by itself,
	where <number> is the number of the philosopher.
*/
// Thanks to @Per Bergqvist for letting me know how could I improve my submission.
package main

import (
	"fmt"
	"sync"
)

var eatingWGroup sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philo struct {
	id             int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

func (p Philo) eat(permissions chan bool, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	// Едим только три раза
	for i := 0; i < 3; i++ {
		<-permissions
		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		fmt.Printf("Философ № %d : 'Ом ном ном!'\n", p.id)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()

		fmt.Printf("Философ № %d : 'Я насытился!'\n", p.id)
		done <- true
	}
}

func host(permissions chan bool, done chan bool) {
	for {
		select {
		case <-done:
			permissions <- true
		}
	}
}

func main() {
	//Создаем вилки
	ChopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	// Приглашаем на обед: Ницше, Фройда, Юнга, Кьеркигора и Марка Аврелия
	Philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philo{
			i,
			ChopSticks[i],       // Левая вилка
			ChopSticks[(i+1)%5], // Правая вилка
		}
	}

	done := make(chan bool)
	permissions := make(chan bool, 2)

	permissions <- true
	permissions <- true

	// Подаем пищу
	for i := 0; i < 5; i++ {
		eatingWGroup.Add(1)
		go Philosophers[i].eat(permissions, done, &eatingWGroup)
	}

	go host(permissions, done)

	eatingWGroup.Wait()
}
