package main

import "fmt"

type CarsFactory interface {
	makeEngine() iEngine
	makeBrakes() iBrakes
}

type iEngine interface {
	setVolume(volume float64)
	setPower(power int)
	getVolume() float64
	getPower() int
}

type iBrakes interface {
	createBrakes(brakes string)
	getBrakes() string
}

type brakes struct {
	name string
}

type engine struct {
	volume float64
	power  int
}

type lambo struct{}

type lamboEngine struct {
	engine
}

type lamboBrakes struct {
	brakes
}

func (f *lambo) makeEngine() iEngine {
	return &lamboEngine{engine{
		volume: 4.0,
		power:  874,
	}}
}

func (f *lambo) makeBrakes() iBrakes {
	return &lamboBrakes{brakes{name: "brembo"}}
}

type fiat struct{}

type fiatEngine struct {
	engine
}

type fiatBrakes struct {
	brakes
}

func (l *fiat) makeEngine() iEngine {
	return &fiatEngine{engine{
		volume: 1.6,
		power:  98,
	}}
}

func (l *fiat) makeBrakes() iBrakes {
	return &fiatBrakes{brakes{name: "baraban"}}
}

func (f *brakes) createBrakes(brakes string) {
	f.name = brakes
}

func (f *brakes) getBrakes() string {
	return f.name
}

func (f *engine) setVolume(volume float64) {
	f.volume = volume
}

func (f *engine) setPower(power int) {
	f.power = power
}

func (f *engine) getVolume() float64 {
	return f.volume
}

func (f *engine) getPower() int {
	return f.power
}

func getCarsFactory(item string) (CarsFactory, error) {
	switch item {
	case "lambo":
		return &lambo{}, nil
	case "fiat":
		return &fiat{}, nil
	}
	return nil, fmt.Errorf("invalid type")
}

func main() {
	lambo, _ := getCarsFactory("lambo")
	//fiat, _ := getCarsFactory("fiat")
	lamboEngine := lambo.makeEngine()
	//fiatEngine := fiat.makeEngine()
	lamboBrakes := lambo.makeBrakes()
	//fiatBrakes := fiat.makeBrakes()
	printDetailsEngine(lamboEngine)
	printDetailsBrakes(lamboBrakes)

}

func printDetailsEngine(e iEngine) {
	fmt.Printf("Volume: %.2f", e.getVolume())
	fmt.Printf("\nPower: %d", e.getPower())
}

func printDetailsBrakes(b iBrakes) {
	fmt.Printf("\nName: %s", b.getBrakes())
}
