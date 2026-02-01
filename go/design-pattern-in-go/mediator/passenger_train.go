package main

import "fmt"

// PassengerTrain は具象Colleague（旅客列車）
type PassengerTrain struct {
	mediator Mediator
}

func NewPassengerTrain(m Mediator) *PassengerTrain {
	return &PassengerTrain{mediator: m}
}

func (t *PassengerTrain) Arrive() {
	if !t.mediator.CanArrive(t) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (t *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	t.mediator.NotifyAboutDeparture()
}

func (t *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted")
	fmt.Println("PassengerTrain: Arrived")
}
