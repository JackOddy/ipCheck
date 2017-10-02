package main

import (
	"fmt"
	"golang.org/x/sync/syncmap"
	"log"
	"sync"
	. "time"
)

var Sites = []string{
	"http://checkip.amazonaws.com/",
	"http://ip.appspot.com/",
	"http://icanhazip.com/",
	"https://api.ipify.com/",
	"http://ip-addr.es/",
}

var PrintAddress = func(address string) {
	fmt.Println("Your address is", address)
}

func startTimer() {
	go func() {
		Sleep(Second * 10)
		log.Fatal("Could not get your address, try again later")
	}()
}

func main() {
	var verifiedAddress string
	foundAddresses := syncmap.Map{}
	addresses := make(chan string, 2)
	workers := []*Worker{}

	wg := sync.WaitGroup{}

	for _, url := range Sites {
		worker := Worker{url, addresses, make(chan bool), make(chan bool)}
		workers = append(workers, &worker)
		wg.Add(1)
		go worker.Start(&wg)
	}

	startTimer()

	for {
		address := <-addresses
		if _, matched := foundAddresses.LoadOrStore(address, true); matched {
			verifiedAddress = address
			break
		}
	}

	for _, worker := range workers {
		go worker.Stop()
	}

	wg.Wait()

	PrintAddress(verifiedAddress)
}
