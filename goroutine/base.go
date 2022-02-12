package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello(i int) {
	defer wg.Done()
	fmt.Printf("value:%d\n", i)
}

