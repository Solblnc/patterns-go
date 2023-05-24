package main

import "fmt"

// Builder - provide a builder interface
type Builder interface {
	makeEngine()
	makeBrakes()
	setColor()
	getCar() car1
}

type car1 struct {
	engine float64
	brakes string
	color  string
}

func getBuilder(carType string) Builder {
	if carType == "sport" {
		return &sportBuilder{}
	}
	return nil
}

type sportBuilder struct {
	engine float64
	brakes string
	color  string
}

func newSportBuilder() *sportBuilder {
	return &sportBuilder{}
}

func (s *sportBuilder) makeEngine() {
	s.engine = 4.2
}

func (s *sportBuilder) makeBrakes() {
	s.brakes = "brembo"
}

func (s *sportBuilder) setColor() {
	s.color = "red"
}

func (s *sportBuilder) getCar() car1 {
	return car1{
		engine: s.engine,
		brakes: s.brakes,
		color:  s.color,
	}
}

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildCar() car1 {
	d.builder.makeEngine()
	d.builder.makeBrakes()
	d.builder.setColor()
	return d.builder.getCar()
}

func main() {
	sportBuilder := getBuilder("sport")
	director := newDirector(sportBuilder)
	sportCar := director.buildCar()

	fmt.Printf("Engine volume: %.2f\n", sportCar.engine)
	fmt.Printf("Brakes: %s\n", sportCar.brakes)
	fmt.Printf("Color %s\n", sportCar.color)

}
