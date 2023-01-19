package http

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Init(c chan string) {
	index, err := os.ReadFile("/Users/ltucillo/Projetos/Pessoal/SimDashboardWeb/server/static/index.html")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(index)
		w.WriteHeader(200)
	})

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		respf, ok := w.(http.Flusher)
		if !ok {
			panic("not flushable")
		}
		w.WriteHeader(200)
		for {
			fmt.Fprint(w, <-c)
			respf.Flush()
			time.Sleep(time.Millisecond * 500)
		}
	})

	http.ListenAndServe(":8080", nil)
}
