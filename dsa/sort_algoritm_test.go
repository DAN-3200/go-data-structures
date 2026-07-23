package dsa_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"app/dsa"
)

func Test_Fibonacci(t *testing.T) {
	uc := new(dsa.Fibo)
	uc.Count = 1
	result := uc.Fibonacci(1, 0)
	require.Equal(t, 4181, result)
}

func Test_BubbleSort(t *testing.T) {
	result := dsa.BubbleSort([]int{8, 7, 3, 10, 28, 4})
	require.Equal(t, []int{3, 4, 7, 8, 10, 28}, result)
}
