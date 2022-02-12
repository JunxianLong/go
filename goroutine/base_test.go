package main

import "testing"

func TestHello(t *testing.T) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Hello(i)
	}
	wg.Wait()
}
