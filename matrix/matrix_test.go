package matrix

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	m := New(10, 5)
	rows, cols := m.Size()
	assert.Equal(10, rows)
	assert.Equal(5, cols)
	assert.Zero(m.Get(0, 0))
	assert.Zero(m.Get(9, 4))
}

func TestIdentitiy(t *testing.T) {
	assert := assert.New(t)

	id := Identity(5)
	rows, cols := id.Size()
	assert.Equal(5, rows)
	assert.Equal(5, cols)
	assert.Equal(1, id.Get(0, 0))
	assert.Equal(1, id.Get(1, 1))
	assert.Equal(1, id.Get(2, 2))
	assert.Equal(1, id.Get(3, 3))
	assert.Equal(1, id.Get(4, 4))
	assert.Equal(0, id.Get(0, 1))
	assert.Equal(0, id.Get(1, 0))
	assert.Equal(0, id.Get(4, 0))
	assert.Equal(0, id.Get(0, 4))
}

func TestNewFromArrays(t *testing.T) {
	assert := assert.New(t)

}

func TestOnes(t *testing.T) {
	assert := assert.New(t)

	ones := Ones(5, 10)
	rows, cols := ones.Size()
	assert.Equal(5, rows)
	assert.Equal(10, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			assert.Equal(1, ones.Get(row, col))
		}
	}
}
