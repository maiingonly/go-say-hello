package go_say_hello

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableOne(t *testing.T) {
	// create struct first
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "maiing",
			request:  "maiing",
			expected: "hello maiing",
		},
		{
			name:     "ucup",
			request:  "ucup",
			expected: "hello ucup",
		},
		{
			name:     "otong",
			request:  "otong",
			expected: "hello otong",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	// subtest Theory 4

	t.Run("ucup", func(t *testing.T) {
		result := SayHello("ucup")
		require.Equal(t, "hello ucup", result)
	})
	t.Run("otong", func(t *testing.T) {
		result := SayHello("otong")
		require.Equal(t, "hello otong", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("mengkoneksikan ke database A (before unit test)")
	// theory 3 >>> before after unit test
	//test main only running in one package
	m.Run() // running all unit test

	fmt.Println("after unit test")
}
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		//theory 2 skip test
		t.Skip("can't run on mac os") // func Skip()
	}
	result := SayHello("maiing")
	require.Equal(t, "hello maiing", result)

}
func TestCalculator(t *testing.T) {

	result := Value{1, 25, 10}.Multiplication()
	// Theory 1
	//assert akan memanggil Fail(), artinya eksekusi unit test akan tetap di lanjutkan
	assert.Equal(t, 250, result, "expect and result not same")

	//require akan memanggil FailNow(), artinya eksekusi unit test tidak akan di lanjutkan kebawahnya
	require.NotEqual(t, 20, result, "they should not be equal")
	fmt.Println("dieksekusi")

}
