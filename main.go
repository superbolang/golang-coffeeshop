package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CoffeeOrder struct {
	Type      string
	Size      string
	Flavoring string
}

var coffeeType = []string{"Hot", "Ice"}
var coffeeSize = []string{"Regular", "Medium", "Large"}
var coffeeFlavor = []string{"Americano", "Latte", "Cappuccino", "Espresso", "Black", "Doppio", "Cortado", "Red Eye"}

// Create random order
func RandomOrder(n int) []CoffeeOrder {
	// Ceate CoffeeOrder slice
	orderResult := []CoffeeOrder{}

	// Create CoffeeOrder instance
	orderCoffe := CoffeeOrder{}

	for range n {
		orderCoffe.Type = coffeeType[rand.Intn(2)]
		orderCoffe.Size = coffeeSize[rand.Intn(3)]
		orderCoffe.Flavoring = coffeeFlavor[rand.Intn(8)]
		orderResult = append(orderResult, orderCoffe)
	}
	return orderResult
}

// original function
func workOrder(b int, q int, o CoffeeOrder) {
	startTime := time.Now()
	fmt.Printf("\nBarista number %d receive order number %d at %v, order: %v\n", b, q, time.Now().Format("02-01-2006 15:04:05.000"), o)
	fmt.Printf("Barista number %d serving order number %d\n", b, q)
	fmt.Printf("Barista number %d finish order number %d\n", b, q)
	fmt.Printf("Order number %d completed in %v\n", q, time.Since(startTime))
	fmt.Println()
}

// function with waitgroup
func workOrderWait(b int, q int, o CoffeeOrder, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Printf("\nBarista number %d receive order number %d at %v, order: %v\n", b, q, time.Now().Format("02-01-2006 15:04:05.000"), o)
	fmt.Printf("Barista number %d serving order number %d\n", b, q)
	fmt.Printf("Barista number %d finish order number %d\n", b, q)
	fmt.Printf("Order number %d completed in %v\n", q, time.Since(startTime))
	fmt.Println()
}

// Create a struct to save channel result
type Result struct {
	Receive  string
	Serving  string
	Finish   string
	Complete string
}

// function with channel
func workOrderChannel(b int, q int, o CoffeeOrder, wg *sync.WaitGroup, result chan<- Result) {
	defer wg.Done()
	startTime := time.Now()
	receive := fmt.Sprintf("\nBarista number %d receive order number %d at %v, order: %v\n", b, q, time.Now().Format(time.StampMilli), o)
	serving := fmt.Sprintf("Barista number %d serving order number %d\n", b, q)
	finish := fmt.Sprintf("Barista number %d finish order number %d\n", b, q)
	complete := fmt.Sprintf("Order number %d completed in %v\n", q, time.Since(startTime))

	result <- Result{Receive: receive, Serving: serving, Finish: finish, Complete: complete} // Sending to channel
}

// Simulate order without goroutine
func orderWithouthGoroutine() {
	// WITHOUT GOROUTINE - SEQUENTIAL
	var number int

	fmt.Println("\n=== ORDER WITHOUT GOROUTINE ===")
	fmt.Print("\n")
	fmt.Print("Number of order: ")
	fmt.Scan(&number)
	incomingOrder := RandomOrder(number)
	start := time.Now()

	for q, o := range incomingOrder {
		q++
		var numberBarista int = rand.Intn(5)
		workOrder(numberBarista, q, o)
	}

	fmt.Printf("Squential process take time: %v\n\n", time.Since(start))
}

// Simulate order with goroutine
func orderWithGoroutine() {
	// WITH GOROUTINE - CONCURRENT
	var number int

	fmt.Println("\n=== ORDER WITH GOROUTINE ===")
	fmt.Print("\n")
	fmt.Print("Number of order: ")
	fmt.Scan(&number)
	incomingOrder := RandomOrder(number)
	start := time.Now()

	for q, o := range incomingOrder {
		q++
		var numberBarista int = rand.Intn(5)
		go workOrder(numberBarista, q, o)
	}

	fmt.Printf("Concurrent process take time: %v\n\n", time.Since(start))

	// The problem with goroutine : when main function goroutine finish faster than schemed goroutine, the process stops and the remaining goroutines are not executed
}

// Simulate goroutine with WaitGroup
func orderWait() {
	var wg sync.WaitGroup

	// WITH WAITGGROUP - CONCURRENT BUT STILL RANDOM
	var number int

	fmt.Println("\n=== ORDER WITH WAITGROUP ===")
	fmt.Print("\n")
	fmt.Print("Number of order: ")
	fmt.Scan(&number)
	incomingOrder := RandomOrder(number)
	start := time.Now()

	for q, o := range incomingOrder {
		wg.Add(1)
		q++
		var numberBarista int = rand.Intn(5)
		go workOrderWait(numberBarista, q, o, &wg)

	}

	wg.Wait()

	fmt.Printf("Concurrent and WaitGroup process take time: %v\n\n", time.Since(start))

}

// Simulate goroutine with channel, the idea is to collect goroutines via channel to make a queue, so one order will be handled by one goroutine from receive to finish
func orderWithChannel() {
	var wg sync.WaitGroup

	// WITH CHANNEL - CONCURRENT AND IN ORDER
	var number int

	fmt.Println("\n=== ORDER WITH WAITGROUP AND CHANNEL ===")
	fmt.Print("\n")
	fmt.Print("Number of order: ")
	fmt.Scan(&number)
	incomingOrder := RandomOrder(number)
	start := time.Now()

	// Create a buffered channel to store/receive result, channel size is as big as number of order
	results := make(chan Result, number)

	for q, o := range incomingOrder {
		wg.Add(1)
		q++
		var numberBarista int = rand.Intn(5)
		go workOrderChannel(numberBarista, q, o, &wg, results)
	}

	// Start a goroutine closure to close channel [IMPORTANT], to terminate `range results` loop
	go func() {
		wg.Wait()
		close(results)
	}()

	// Iterate over orders to print results
	for i := 0; i < len(incomingOrder); i++ {
		receiveResult := <-results // receive from channel
		fmt.Println(receiveResult)
	}

	fmt.Printf("\nWaitGroup via channel process take time: %v\n\n", time.Since(start))

}

func main() {
	// Comment/uncomment for execution
	// orderWithouthGoroutine()
	orderWithGoroutine()
	// orderWait()
	// orderWithChannel()
}
