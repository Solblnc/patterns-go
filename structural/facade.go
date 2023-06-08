package main

import "fmt"

type carFacade struct {
	engine       *engine
	brakes       *brakes
	transmission *transmission
}

func newCarFacade() *carFacade {
	fmt.Println("Starting create car")
	carFacade1 := &carFacade{
		engine:       newEngine(),
		brakes:       newBrakes(),
		transmission: newTransmission(),
	}
	fmt.Println("Car created")
	return carFacade1
}

func newTransmission() *transmission {
	return &transmission{brand: "ZF"}
}

func newBrakes() *brakes {
	return &brakes{name: "brembo"}
}

func newEngine() *engine {
	return &engine{power: 710}
}

type engine struct {
	power int
}

type brakes struct {
	name string
}

type transmission struct {
	brand string
}

func main() {
	car := newCarFacade()

	fmt.Println(*car.engine)
	fmt.Println(*car.brakes)
	fmt.Println(*car.transmission)
}
