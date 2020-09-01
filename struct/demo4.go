package main

import (
	"encoding/json"
	"fmt"
)

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,omitempty"`
}

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = false
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
