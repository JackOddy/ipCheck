package main

import (
	. "ipCheck/test"
	"testing"
)

func setUpServers(t *testing.T, urls []string) (finish func()) {
	ready := make(chan func())

	go TestServer(t, urls, ready)
	finish = <-ready
	return
}

func TestIpCheck(t *testing.T) {
	Sites = []string{
		"http://localhost:3000/1", // wil fail
		"http://localhost:4000/2", //will fail
		"http://localhost:5000",
		"http://localhost:6000",
	}
	Ports := []string{
		":3000",
		":4000",
		":5000",
		":6000",
	}

	defer setUpServers(t, Ports)()
	expected := "12345.12345\n"

	var result string
	PrintAddress = func(r string) {
		result = r
	}

	main()

	if result != expected {
		t.Fatal("Expected", expected, "to equal", result)
	}
}
