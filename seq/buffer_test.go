package seq

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestBuffer(t *testing.T) {
	assert := assert.New(t)

	buffer := NewBuffer()

	buffer.Enqueue(1)
	assert.Equal(1, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(1, buffer.PeekBack())

	buffer.Enqueue(2)
	assert.Equal(2, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(2, buffer.PeekBack())

	buffer.Enqueue(3)
	assert.Equal(3, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(3, buffer.PeekBack())

	buffer.Enqueue(4)
	assert.Equal(4, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(4, buffer.PeekBack())

	buffer.Enqueue(5)
	assert.Equal(5, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(5, buffer.PeekBack())

	buffer.Enqueue(6)
	assert.Equal(6, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(6, buffer.PeekBack())

	buffer.Enqueue(7)
	assert.Equal(7, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(7, buffer.PeekBack())

	buffer.Enqueue(8)
	assert.Equal(8, buffer.Len())
	assert.Equal(1, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value := buffer.Dequeue()
	assert.Equal(1, value)
	assert.Equal(7, buffer.Len())
	assert.Equal(2, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(2, value)
	assert.Equal(6, buffer.Len())
	assert.Equal(3, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(3, value)
	assert.Equal(5, buffer.Len())
	assert.Equal(4, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(4, value)
	assert.Equal(4, buffer.Len())
	assert.Equal(5, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(5, value)
	assert.Equal(3, buffer.Len())
	assert.Equal(6, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(6, value)
	assert.Equal(2, buffer.Len())
	assert.Equal(7, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(7, value)
	assert.Equal(1, buffer.Len())
	assert.Equal(8, buffer.Peek())
	assert.Equal(8, buffer.PeekBack())

	value = buffer.Dequeue()
	assert.Equal(8, value)
	assert.Equal(0, buffer.Len())
	assert.Zero(buffer.Peek())
	assert.Zero(buffer.PeekBack())
}

func TestBufferClear(t *testing.T) {
	assert := assert.New(t)

	buffer := NewBuffer()
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)
	buffer.Enqueue(1)

	assert.Equal(8, buffer.Len())

	buffer.Clear()
	assert.Equal(0, buffer.Len())
	assert.Zero(buffer.Peek())
	assert.Zero(buffer.PeekBack())
}

func TestBufferArray(t *testing.T) {
	assert := assert.New(t)

	buffer := NewBuffer()
	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)
	buffer.Enqueue(4)
	buffer.Enqueue(5)

	contents := buffer.Array()
	assert.Len(5, contents)
	assert.Equal(1, contents[0])
	assert.Equal(2, contents[1])
	assert.Equal(3, contents[2])
	assert.Equal(4, contents[3])
	assert.Equal(5, contents[4])
}

func TestBufferEach(t *testing.T) {
	assert := assert.New(t)

	buffer := NewBuffer()

	for x := 1; x < 17; x++ {
		buffer.Enqueue(float64(x))
	}

	called := 0
	buffer.Each(func(_ int, v float64) {
		if v == float64(called+1) {
			called++
		}
	})

	assert.Equal(16, called)
}

func TestNewBuffer(t *testing.T) {
	assert := assert.New(t)

	empty := NewBuffer()
	assert.NotNil(empty)
	assert.Zero(empty.Len())
	assert.Equal(bufferDefaultCapacity, empty.Capacity())
	assert.Zero(empty.Peek())
	assert.Zero(empty.PeekBack())
}

func TestNewBufferWithValues(t *testing.T) {
	assert := assert.New(t)

	values := NewBuffer(1, 2, 3, 4, 5)
	assert.NotNil(values)
	assert.Equal(5, values.Len())
	assert.Equal(1, values.Peek())
	assert.Equal(5, values.PeekBack())
}

func TestBufferGrowth(t *testing.T) {
	assert := assert.New(t)

	values := NewBuffer(1, 2, 3, 4, 5)
	for i := 0; i < 1<<10; i++ {
		values.Enqueue(float64(i))
	}

	assert.Equal(1<<10-1, values.PeekBack())
}
