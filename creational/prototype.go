package main

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type bike struct {
	name string
}

func (b *bike) print(ident string) {
	fmt.Println(ident + b.name)
}

func (b *bike) clone() inode {
	return &bike{name: b.name + "_clone"}
}

type garage struct {
	bikes []inode
	name  string
}

func (b *garage) print(indent string) {
	fmt.Println(indent + b.name)
	for _, i := range b.bikes {
		i.print(indent + indent)
	}
}

func (b *garage) clone() inode {
	cloneBikes := &garage{name: b.name + "_clone"}
	var tempBikes []inode
	for _, i := range b.bikes {
		copy := i.clone()
		tempBikes = append(tempBikes, copy)
	}
	cloneBikes.bikes = tempBikes
	return cloneBikes
}

func main() {
	bike1 := &bike{name: "bmw"}
	bike2 := &bike{name: "ktm"}
	bike3 := &bike{name: "honda"}

	garage1 := &garage{
		bikes: []inode{bike1},
		name:  "Garage 1",
	}

	garage2 := &garage{
		bikes: []inode{garage1, bike2, bike3},
		name:  "Garage 2",
	}

	fmt.Println("\nPrinting Garage 1")
	garage2.print("   ")

	cloneGarage := garage2.clone()
	fmt.Println("\nPrinting Garage 2")
	cloneGarage.print("   ")
}
