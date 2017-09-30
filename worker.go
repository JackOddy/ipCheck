package main

import (
	"fmt"
	"net/http"
)

var Report = func(s string, channel chan string) {
	channel <- s
}

type Worker struct {
	url, addr string
	channel   chan string
}

func (self Worker) ping() (resp *http.Response, ok bool) {
	resp, err := http.Get(self.url)

	if err != nil {
		ok = false
	}

	return
}

func (self Worker) report() {
	go Report(self.addr, self.channel)
}
