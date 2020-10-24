package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello type
type Hello struct {
	l *log.Logger
}

// NewHello func
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Bismillah bisa")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Sorry guys, maybe next time.", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Assalamu'alaikum %s", d)
}
