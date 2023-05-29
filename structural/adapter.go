package main

import "fmt"

type car interface {
	insertFastEngine()
}

type ferrari struct {
}

func (f *ferrari) insertFastEngine() {
	fmt.Println("Insert fast engine in ferrari")
}

type lada struct {
}

func (l *lada) insertSlowEngine() {
	fmt.Println("Insert slow engine in lada")
}

type ladaAdapter struct {
	ladaEngine *lada
}

func (l *ladaAdapter) insertFastEngine() {
	l.ladaEngine.insertSlowEngine()
}

type client struct {
}

func (c *client) insertFastEngineInCar(car car) {
	car.insertFastEngine()
}

func main() {
	client := &client{}
	ferrari := &ferrari{}
	client.insertFastEngineInCar(ferrari)
	lada := &lada{}
	ladaAdapter := &ladaAdapter{ladaEngine: lada}
	client.insertFastEngineInCar(ladaAdapter)
}
