package main

import (
	"SimDashboardWeb/server/http"
	"SimDashboardWeb/server/tcp"
)

func main() {
	ch := make(chan string)
	go tcp.CreateServer(ch)

	http.Init(ch)
}
