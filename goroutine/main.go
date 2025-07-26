package main

import (
	"fmt"
	"sync"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task2")

}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task3")
}

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}
func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go task1(&wg)
	wg.Add(1)
	go task2(&wg)
	wg.Add(1)
	go task3(&wg)

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main Completed")

}
