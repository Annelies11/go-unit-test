package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Testing")
	m.Run()
	fmt.Println("After Testing")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run in Windows")
	}

	result := SayHello("Aris")
	require.Equal(t, "Hello Aris", result, "Result must be 'Hello Aris'")
}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("Aris")
	require.Equal(t, "Hello Aris", result, "Harusnya 'Hello Aris' blok")
	fmt.Println("Lha iki gak ketoro")
}

func TestSayHelloAssert(t *testing.T) {
	result := SayHello("Aris")
	assert.Equal(t, "Hello Aris", result, "Harusnya 'Hello Aris' blok")
	fmt.Println("Test with Assert done")
}

func TestSayHello(t *testing.T) {
	res := SayHello("Aris")

	if res != "Hello Aris" {
		t.Fatal("Harusnya Hai Aris")
	}
	fmt.Println("TestSayHello Done")
}

func TestSayHelloBro(t *testing.T) {
	res := SayHello("Bro")

	if res != "Hello Bro" {
		t.FailNow()
	}
	fmt.Println("TestSayHelloBro Done")
}
