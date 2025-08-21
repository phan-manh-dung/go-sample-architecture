package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// 1. Cách 1
	var (
		input  = 1 // dữ liệu đưa vào
		output = 2 // kết quả mong muốn (khác thì fail)
	)

	actual := AddOne(input)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}

	// 2. Cách 2 dùng assert
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should return 3") // 3 là kết quả mong muốn
	assert.NotEqual(t, AddOne(2), 4, "AddOne(2) should not return 4")
	// Nhiều method assert khác
}
