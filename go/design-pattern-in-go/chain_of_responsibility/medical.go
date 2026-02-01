package main

import "fmt"

// Medical は薬局部門（具象ハンドラー）
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	if m.next != nil {
		m.next.execute(p)
	}
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
