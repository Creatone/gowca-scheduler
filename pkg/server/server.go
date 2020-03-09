package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Host string
	Port int
}

func (s Server) Serve() {
	fmt.Println(s.Host)
	fmt.Println(s.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	ip := fmt.Sprintf("%s:%d", s.Host, s.Port)
	http.ListenAndServe(ip, nil)
}
