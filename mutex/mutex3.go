package main

import (
	"fmt"
	"sync"
)

// Alice 的账户余额
var balance int
var sema = make(chan struct{}, 1)

// 存钱
func Deposit(amount int) {
	//sema <- struct{}{}
	balance = balance + amount
	//<-sema
}

// 查看余额
func Balance() int {
	return balance
}
func alice(amount int) {
	fmt.Printf("Alice deposit $%d\n", amount)
	Deposit(amount)
}

func bob(amount int) {
	fmt.Printf("Bob deposit $%d\n", amount)
	Deposit(amount)
}
func main() {
	//单 goroutine 顺序存取
	/*alice(200);
	bob(100);
	fmt.Printf("balance: $%d\n", bank.Balance()) // balance: 300*/

	/*var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		alice(100)
		wg.Done()
	}()
	go func() {
		bob(200)
		wg.Done()
	}()

	wg.Wait()*/
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			alice(100)
			wg.Done()
		}()
		go func() {
			bob(200)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("balance: $%d\n", Balance())
}
