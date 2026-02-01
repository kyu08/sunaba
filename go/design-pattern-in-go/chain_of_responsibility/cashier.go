package main

import "fmt"

// Cashier は会計部門（具象ハンドラー、連鎖の終端）
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
		return
	}
	fmt.Println("Cashier getting money from patient")
	p.paymentDone = true
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
