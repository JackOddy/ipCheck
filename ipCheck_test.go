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
		":3000",
		":4000",
		":5000",
		":6000",
	}

	defer setUpServers(t, Sites)()
	expected := "testIp"

	var result string
	Report = func(r string) {
		result = r
	}

	main()

	if result != expected {
		t.Fatal("Expected", expected, "to equal", result)
	}

}
