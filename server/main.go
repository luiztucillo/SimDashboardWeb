package main

import (
	"SimDashboardWeb/server/tcp"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func receive(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan string)
	go tcp.CreateServer(ch)
	go receive(ch)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Fprintf(w, string(p[:n]))
	}
}
