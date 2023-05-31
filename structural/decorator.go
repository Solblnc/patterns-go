package main

import "fmt"

type icar interface {
	getPower() int
}

type nissan struct {
}

func (n *nissan) getPower() int {
	return 140
}

type lambo struct {
}

func (l *lambo) getPower() int {
	return 720
}

type exhaust struct {
	car icar
}

func (e *exhaust) getPower() int {
	carPower := e.car.getPower()
	return carPower + 25
}

type turbine struct {
	car icar
}

func (t *turbine) getPower() int {
	carPower := t.car.getPower()
	return carPower + 70
}

func main() {
	nissanCar := &nissan{}

	nissanExhaust := &exhaust{car: nissanCar}
	fmt.Printf("Power of nissan with sport exhaust: %d\n", nissanExhaust.getPower())

	lamboCar := &lambo{}
	lamboTurbine := &turbine{car: lamboCar}
	fmt.Printf("Power of lambo with turbine: %d\n", lamboTurbine.getPower())

}
