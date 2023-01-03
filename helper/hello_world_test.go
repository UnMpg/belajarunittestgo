package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkSub(b *testing.B) {
	b.Run("Fendy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fendy")
		}
	})

	b.Run("Erid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Erid")
		}
	})

}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Eko",
			Request: "Eko",
		}, {
			Name:    "Fendy",
			Request: "Fendy",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")

	assert.Equal(t, "Hello Eko", result, "resul must be Hello Eko")
	fmt.Println("Test assert with done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")

	require.Equal(t, "Hello Eko", result, "resul must be Hello Eko")
	fmt.Println("Test assert with done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		t.Error("result Must be  Hello Eko")
	}

	fmt.Println("test hello world eko done")
}

func TestHelloWorldFendy(t *testing.T) {
	result := HelloWorld("Fendy")

	if result != "Hello Fendy" {
		t.Fatal("Result must be Hello Fendy")
	}

	fmt.Println("test hello world fendy done")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("before unit test")

	m.Run()

	//after

	fmt.Println("after unit test")
}

// func TestSkip(t *testing.T) {
// 	if runtime.GOOS == "windows" {
// 		t.Skip("can not run on Windoes")
// 	}
// 	result := HelloWorld("Fendy")
// 	assert.Equal(t, "Hello Fendy", result, "resul must be Hello Eko")
// }

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")

		assert.Equal(t, "Hello Eko", result, "resul must be Hello Eko")
	})

	t.Run("fendy", func(t *testing.T) {
		result := HelloWorld("Fendy")

		assert.Equal(t, "Hello Fendy", result, "resul must be Hello Fendy")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fendy",
			request:  "Fendy",
			expected: "Hello Fendy",
		}, {
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		}, {
			name:     "Erik",
			request:  "Erik",
			expected: "Hello Erik",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
