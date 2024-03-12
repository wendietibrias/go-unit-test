package helper

import (
	"fmt"
	"testing"

	"runtime"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSubSayHelloWorld(b *testing.B) {

	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "Wendi",
			request: "Parker",
		},
		{
			name:    "Randi",
			request: "Randi",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHelloWorld(test.request)
			}
		})
	}

}

func TestSubTest(t *testing.T) {
	tests := []struct {
		expected string
		name     string
		request  string
	}{
		{
			name:     "SayHelloWorld(Randi)",
			expected: "Hello World for you Randi",
			request:  "Randi",
		},
		{
			name:     "SayHelloWorld(Wendi)",
			expected: "Hello World for you Randi",
			request:  "Wendi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Must be equal")
		})
	}
}

// func TestMain(m *testing.M) {
// 	fmt.Println("Before testing")

// 	m.Run()

// 	fmt.Println("After testing")
// }

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "Wendi" {
		t.Skip("Tidak dapat menjalankan unit testing")
	}
	result := SayHelloWorld("Randi")
	assert.Equal(t, "Hello World for you Wendi", result, "Must be equal")
}

func TestSayHelloWorld(t *testing.T) {
	result := SayHelloWorld("Randi")
	assert.Equal(t, "Hello World for you Wendi", result, "Must be equal")
}
