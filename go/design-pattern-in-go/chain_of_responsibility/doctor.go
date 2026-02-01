package main

import "fmt"

// Doctor は医師部門（具象ハンドラー）
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
