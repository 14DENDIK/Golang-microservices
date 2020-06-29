package handlers

import (
  "log"
  "bitbucket.org/14DENDIK/microservices/data"
  "net/http"
  // "encoding/json"

)

type Products struct {
  l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
  return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    p.getProducts(rw, r)
    return
  }

  rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
  lp := data.GetProducts()
  err := lp.ToJSON(rw)
  if err != nil {
    http.Error(rw, "Cannot marshal passed data", http.StatusInternalServerError)
  }
}
