package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

// once - protect us from race condition
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	inst := GetInstance()
	fmt.Println(inst)
}
