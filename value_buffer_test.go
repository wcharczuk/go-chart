package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRingBuffer(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()

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
	assert.Nil(buffer.Peek())
	assert.Nil(buffer.PeekBack())
}

func TestRingBufferClear(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()
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
	assert.Nil(buffer.Peek())
	assert.Nil(buffer.PeekBack())
}

func TestRingBufferContents(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()
	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)
	buffer.Enqueue(4)
	buffer.Enqueue(5)

	contents := buffer.Contents()
	assert.Len(contents, 5)
	assert.Equal(1, contents[0])
	assert.Equal(2, contents[1])
	assert.Equal(3, contents[2])
	assert.Equal(4, contents[3])
	assert.Equal(5, contents[4])
}

func TestRingBufferDrain(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()
	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)
	buffer.Enqueue(4)
	buffer.Enqueue(5)

	contents := buffer.Drain()
	assert.Len(contents, 5)
	assert.Equal(1, contents[0])
	assert.Equal(2, contents[1])
	assert.Equal(3, contents[2])
	assert.Equal(4, contents[3])
	assert.Equal(5, contents[4])

	assert.Equal(0, buffer.Len())
	assert.Nil(buffer.Peek())
	assert.Nil(buffer.PeekBack())
}

func TestRingBufferEach(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()

	for x := 1; x < 17; x++ {
		buffer.Enqueue(x)
	}

	called := 0
	buffer.Each(func(v interface{}) {
		if typed, isTyped := v.(int); isTyped {
			if typed == (called + 1) {
				called++
			}
		}
	})

	assert.Equal(16, called)
}

func TestRingBufferEachUntil(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()

	for x := 1; x < 17; x++ {
		buffer.Enqueue(x)
	}

	called := 0
	buffer.EachUntil(func(v interface{}) bool {
		if typed, isTyped := v.(int); isTyped {
			if typed > 10 {
				return false
			}
			if typed == (called + 1) {
				called++
			}
		}
		return true
	})

	assert.Equal(10, called)
}

func TestRingBufferReverseEachUntil(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBufferWithCapacity(32)

	for x := 1; x < 17; x++ {
		buffer.Enqueue(x)
	}

	var values []int
	buffer.ReverseEachUntil(func(v interface{}) bool {
		if typed, isTyped := v.(int); isTyped {
			if typed < 10 {
				return false
			}
			values = append(values, typed)
			return true
		}
		panic("value is not an integer")
	})

	assert.Len(values, 7)
	assert.Equal(16, values[0])
	assert.Equal(10, values[6])
}

func TestRingBufferReverseEachUntilUndersized(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()

	for x := 1; x < 17; x++ {
		buffer.Enqueue(x)
	}

	var values []int
	buffer.ReverseEachUntil(func(v interface{}) bool {
		if typed, isTyped := v.(int); isTyped {
			if typed < 10 {
				return false
			}
			values = append(values, typed)
			return true
		}
		panic("value is not an integer")
	})

	assert.Len(values, 7)
	assert.Equal(16, values[0])
	assert.Equal(10, values[6])
}

func TestRingBufferConsume(t *testing.T) {
	assert := assert.New(t)

	buffer := NewRingBuffer()

	for x := 1; x < 17; x++ {
		buffer.Enqueue(x)
	}

	assert.Equal(16, buffer.Len())

	var called int
	buffer.Consume(func(v interface{}) {
		if _, isTyped := v.(int); isTyped {
			called++
		}
	})

	assert.Equal(16, called)
	assert.Zero(buffer.Len())
}
