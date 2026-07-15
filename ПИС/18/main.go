package main

import (
	"fmt"
	"sync"
	"time"
)

func loadUsers(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Load users")
	time.Sleep(2 * time.Second)
	fmt.Println("Users loaded")
}

func loadOrders(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Load orders")
	time.Sleep(3 * time.Second)
	fmt.Println("Orders loaded")
}

func loadProducts(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Load products")
	time.Sleep(1 * time.Second)
	fmt.Println("Products loaded")
}

func main() {

	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(3)

	go loadUsers(&wg)
	go loadOrders(&wg)
	go loadProducts(&wg)

	wg.Wait()
	fmt.Println("End")
}
