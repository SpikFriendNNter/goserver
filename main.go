package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		io.WriteString(w, addr[:strings.Index(addr, ":")])
	})

	http.ListenAndServe(":80", nil)
}
