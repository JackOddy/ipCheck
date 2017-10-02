package main

import (
	"io/ioutil"
	"net/http"
	"sync"
	. "time"
)

type Worker struct {
	Url       string
	Addresses chan string
	Quit      chan bool
	Retry     chan bool
}

func (self *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case <-self.Quit:
				return
			case <-self.Retry:
				go self.Search()
			}
		}

	}()

	self.Search()
}

func (self *Worker) Stop() {
	self.Quit <- true
}

func (self *Worker) Search() {
	if address, ok := self.getAddress(); ok {
		self.Addresses <- address
		return
	}
	Sleep(Second)
	self.Retry <- true
}

func (self *Worker) getAddress() (address string, ok bool) {
	resp, err := http.Get(self.Url)

	if err != nil {
		return
	}

	if resp.StatusCode > 399 {
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return string(bytes), true
}
