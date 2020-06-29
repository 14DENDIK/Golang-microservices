package handlers

import (
  "net/http"
  "log"
  // "io/ioutil"
  // "fmt"
)

type Goodbye struct {
  l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
  return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  g.l.Printf("Goodbye handler")
  rw.Write([]byte("Byee page"))
}
