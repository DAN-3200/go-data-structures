// AAA (Arrange, Act, Assert)
package dsa_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/dsa"
)

func Test_Fibonacci(t *testing.T) {
	// arrange
	uc := new(dsa.Fibo)
	uc.Count = 1
	
	// act
	result := uc.Fibonacci(1, 0)
	
	// assert
	assert.Equal(t, 4181, result)
}

func Test_BubbleSort(t *testing.T) {
	// arrange
	list:= []int{8, 7, 3, 10, 28, 4}
	
	// act
	result := dsa.BubbleSort(list)

	// assert
	assert.Equal(t, []int{3, 4, 7, 8, 10, 28}, result)
}
