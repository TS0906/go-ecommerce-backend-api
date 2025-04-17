package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 3
	)
	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}
}
func TestRequire(t *testing.T) {
	require.Equal(t, 2, AddOne(1))
	fmt.Println("Not executing")
}

// chan thuc thi cau lenh duoi neu test fail
func TestAssert(t *testing.T) {
	assert.Equal(t, 2, AddOne(3))
	fmt.Println("executing")
}