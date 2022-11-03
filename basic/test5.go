package test

import (
	"fmt"
	"time"
)

func Main52() {
	go doSomething("Calculation")
	time.Sleep(time.Second)
}

func doSomething(w string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, w)
	}
}

func Main51() {
	c1 := make(chan string)
	go func() { c1 <- "My Problem1" }()
	fmt.Println(<-c1)
}
