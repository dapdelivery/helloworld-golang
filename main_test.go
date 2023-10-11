package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if  "Hello World!!" != "Hello World!!" {
		t.Fatal("Test fail")
	}
}
