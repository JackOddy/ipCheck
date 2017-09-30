package testServer

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T, urls []string, ready chan func()) {
	var testServers []*http.Server

	http.Handle("/",
		http.FileServer(
			http.Dir("./test/html"),
		),
	)

	for _, url := range urls {
		server := http.Server{Addr: url}
		go server.ListenAndServe()

		testServers = append(testServers, &server)
	}

	ready <- func() {
		for _, server := range testServers {
			if err := server.Shutdown(nil); err != nil {
				panic(err)
			}
		}
	}
}
