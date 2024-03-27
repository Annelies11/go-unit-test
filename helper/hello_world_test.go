package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	bs := []struct {
		name string
		req  string
	}{
		{
			name: "Aris",
			req:  "Aris",
		},
		{
			name: "Mahmudi",
			req:  "Mahmudi",
		},
	}
	for _, bench := range bs {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(bench.req)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Aris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Aris")
		}
	})
	b.Run("Mahmudi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Mahmudi")
		}
	})
}

func BenchmarkSayHelloAris(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Aris")
	}
}

func BenchmarkSayHelloMahmudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Mahmudi")
	}
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		req      string
		expected string
	}{
		{
			name:     "Aris",
			req:      "Aris",
			expected: "Hello Aris",
		},
		{
			name:     "Mahmudi",
			req:      "Mahmudi",
			expected: "Hello Mahmudi",
		},
		{
			name:     "Mavis",
			req:      "Mavis",
			expected: "Hello Mavis",
		},
		{
			name:     "Last",
			req:      "Last",
			expected: "Hello Last",
		},
		{
			name:     "Child",
			req:      "Child",
			expected: "Hello Child",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := SayHello(test.req)
			require.Equal(t, test.expected, res)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Aris", func(t *testing.T) {
		result := SayHello("Aris")
		require.Equal(t, "Hello Aris", result, "Harusnya 'Hello Aris' blok")
	})
	t.Run("Mahmudi", func(t *testing.T) {
		result := SayHello("Mahmudi")
		require.Equal(t, "Hi Mahmudi", result, "Harusnya 'Hello Mahmudi' blok")
	})
}

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
