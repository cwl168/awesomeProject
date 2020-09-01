package main

import (
	"encoding/json"
	"fmt"
)

//{"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}
// Product 商品信息
//type Product struct {
//	Name      string
//	ProductID int64
//	Number    int
//	Price     float64
//	IsOnSale  bool
//}
//{"name":"Xiao mi 6","number":10000,"price":2499,"is_on_sale":"true"}
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"-"` // 表示不进行序列化
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}

//结果
//{"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}
