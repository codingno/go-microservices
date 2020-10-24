package handlers

import (
	"log"
	"net/http"
)

// Goodbye type
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye func
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee guys"))
}
