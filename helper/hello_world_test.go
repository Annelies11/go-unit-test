package helper

import "testing"

func TestSayHello(t *testing.T) {
	res := SayHello("Aris")

	if res != "Hello Aris" {
		panic("Hoi salah oi")
	}
}

func TestSayHelloBro(t *testing.T) {
	res := SayHello("Bro")

	if res != "Hello Bro" {
		panic("Hoi salah oi")
	}
}
