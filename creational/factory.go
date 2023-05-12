package main

import "fmt"

// Car - provides a factory interface
type Car interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
}

// car - defines
type car struct {
	name  string
	power int
}

// Methods of interface
func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setPower(power int) {
	c.power = power
}

func (c *car) getPower() int {
	return c.power
}

type ferrari struct {
	car
}

func newFerrari() Car {
	return &ferrari{
		car{
			name:  "ferrari",
			power: 986,
		}}

}

type lada struct {
	car
}

func newLada() Car {
	return &lada{car{
		name:  "lada",
		power: 98,
	}}
}

// getCar - a Factory method
func getCar(carType string) (Car, error) {
	if carType == "ferrari" {
		return newFerrari(), nil
	}

	if carType == "lada" {
		return newLada(), nil
	}
	return nil, fmt.Errorf("wrong type passed")
}

func main() {
	ferrari, _ := getCar("ferrari")
	lada, _ := getCar("lada")
	printDetails(ferrari)
	printDetails(lada)

}

func printDetails(c Car) {
	fmt.Printf("\nCar: %s ", c.getName())
	fmt.Printf("Power: %d ", c.getPower())

}
