package data

import (
  "time"
  "encoding/json"
  "io"
)

type Product struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Price float32 `json:"price"`
  SKU string `json:sku`
  CreatedOn string `json:"-"`
  UpdatedOn string `json:"-"`
  DeletedOn string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
  e := json.NewEncoder(w)
  return e.Encode(p)
}

func GetProducts() Products {
  return productList
}

var productList = Products {
  &Product {
    ID: 1,
    Name: "Lotte",
    Description: "Coffee with frosty",
    Price: 32.43,
    SKU: "abc312",
    CreatedOn: time.Now().UTC().String(),
    UpdatedOn: time.Now().UTC().String(),
  },
  &Product {
    ID: 2,
    Name: "Espresso",
    Description: "Strong Small Coffee",
    Price: 5.43,
    SKU: "aaa312",
    CreatedOn: time.Now().UTC().String(),
    UpdatedOn: time.Now().UTC().String(),
  },
}
