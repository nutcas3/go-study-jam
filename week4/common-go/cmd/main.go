package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: prime_finder <start> <end>")
		return
	}
	start, err1 := strconv.Atoi(os.Args[1])
	end, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || start > end {
		fmt.Println("Invalid range.")
		return
	}
	var wg sync.WaitGroup
	primes := make(chan int, end-start+1)
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if isPrime(n) {
				primes <- n
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(primes)
	}()
	fmt.Printf("Primes between %d and %d:\n", start, end)
	for p := range primes {
		fmt.Println(p)
	}
}
