package main

import "fmt"

// FreightTrain は具象Colleague（貨物列車）
type FreightTrain struct {
	mediator Mediator
}

func NewFreightTrain(m Mediator) *FreightTrain {
	return &FreightTrain{mediator: m}
}

func (t *FreightTrain) Arrive() {
	if !t.mediator.CanArrive(t) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (t *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	t.mediator.NotifyAboutDeparture()
}

func (t *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	fmt.Println("FreightTrain: Arrived")
}
