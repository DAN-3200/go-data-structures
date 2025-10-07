package module_test

import (
	"app/module"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BubbleSort(t *testing.T) {
	result := module.BubbleSort([]int{8, 7, 3, 10, 28, 4})
	require.Equal(t, result, []int{3, 4, 7, 8, 10, 28})
}
