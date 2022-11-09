package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient " + p.name)
	p.paymentDone = true
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
