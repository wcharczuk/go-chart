package seq

import (
	"fmt"
	"strings"

	util "github.com/wcharczuk/go-chart/util"
)

const (
	bufferMinimumGrow     = 4
	bufferShrinkThreshold = 32
	bufferGrowFactor      = 200
	bufferDefaultCapacity = 4
)

var (
	emptyArray = make([]float64, 0)
)

// NewBuffer creates a new value buffer with an optional set of values.
func NewBuffer(values ...float64) *Buffer {
	var tail int
	array := make([]float64, util.Math.MaxInt(len(values), bufferDefaultCapacity))
	if len(values) > 0 {
		copy(array, values)
		tail = len(values)
	}
	return &Buffer{
		array: array,
		head:  0,
		tail:  tail,
		size:  len(values),
	}
}

// NewBufferWithCapacity creates a new ValueBuffer pre-allocated with the given capacity.
func NewBufferWithCapacity(capacity int) *Buffer {
	return &Buffer{
		array: make([]float64, capacity),
		head:  0,
		tail:  0,
		size:  0,
	}
}

// Buffer is a fifo datastructure that is backed by a pre-allocated array.
// Instead of allocating a whole new node object for each element, array elements are re-used (which saves GC churn).
// Enqueue can be O(n), Dequeue is generally O(1).
// Buffer implements `seq.Provider`
type Buffer struct {
	array []float64
	head  int
	tail  int
	size  int
}

// Len returns the length of the Buffer (as it is currently populated).
// Actual memory footprint may be different.
func (b *Buffer) Len() int {
	return b.size
}

// GetValue implements seq provider.
func (b *Buffer) GetValue(index int) float64 {
	effectiveIndex := (b.head + index) % len(b.array)
	return b.array[effectiveIndex]
}

// Capacity returns the total size of the Buffer, including empty elements.
func (b *Buffer) Capacity() int {
	return len(b.array)
}

// SetCapacity sets the capacity of the Buffer.
func (b *Buffer) SetCapacity(capacity int) {
	newArray := make([]float64, capacity)
	if b.size > 0 {
		if b.head < b.tail {
			arrayCopy(b.array, b.head, newArray, 0, b.size)
		} else {
			arrayCopy(b.array, b.head, newArray, 0, len(b.array)-b.head)
			arrayCopy(b.array, 0, newArray, len(b.array)-b.head, b.tail)
		}
	}
	b.array = newArray
	b.head = 0
	if b.size == capacity {
		b.tail = 0
	} else {
		b.tail = b.size
	}
}

// Clear removes all objects from the Buffer.
func (b *Buffer) Clear() {
	b.array = make([]float64, bufferDefaultCapacity)
	b.head = 0
	b.tail = 0
	b.size = 0
}

// Enqueue adds an element to the "back" of the Buffer.
func (b *Buffer) Enqueue(value float64) {
	if b.size == len(b.array) {
		newCapacity := int(len(b.array) * int(bufferGrowFactor/100))
		if newCapacity < (len(b.array) + bufferMinimumGrow) {
			newCapacity = len(b.array) + bufferMinimumGrow
		}
		b.SetCapacity(newCapacity)
	}

	b.array[b.tail] = value
	b.tail = (b.tail + 1) % len(b.array)
	b.size++
}

// Dequeue removes the first element from the RingBuffer.
func (b *Buffer) Dequeue() float64 {
	if b.size == 0 {
		return 0
	}

	removed := b.array[b.head]
	b.head = (b.head + 1) % len(b.array)
	b.size--
	return removed
}

// Peek returns but does not remove the first element.
func (b *Buffer) Peek() float64 {
	if b.size == 0 {
		return 0
	}
	return b.array[b.head]
}

// PeekBack returns but does not remove the last element.
func (b *Buffer) PeekBack() float64 {
	if b.size == 0 {
		return 0
	}
	if b.tail == 0 {
		return b.array[len(b.array)-1]
	}
	return b.array[b.tail-1]
}

// TrimExcess resizes the capacity of the buffer to better fit the contents.
func (b *Buffer) TrimExcess() {
	threshold := float64(len(b.array)) * 0.9
	if b.size < int(threshold) {
		b.SetCapacity(b.size)
	}
}

// Array returns the ring buffer, in order, as an array.
func (b *Buffer) Array() Array {
	newArray := make([]float64, b.size)

	if b.size == 0 {
		return newArray
	}

	if b.head < b.tail {
		arrayCopy(b.array, b.head, newArray, 0, b.size)
	} else {
		arrayCopy(b.array, b.head, newArray, 0, len(b.array)-b.head)
		arrayCopy(b.array, 0, newArray, len(b.array)-b.head, b.tail)
	}

	return Array(newArray)
}

// Each calls the consumer for each element in the buffer.
func (b *Buffer) Each(mapfn func(int, float64)) {
	if b.size == 0 {
		return
	}

	var index int
	if b.head < b.tail {
		for cursor := b.head; cursor < b.tail; cursor++ {
			mapfn(index, b.array[cursor])
			index++
		}
	} else {
		for cursor := b.head; cursor < len(b.array); cursor++ {
			mapfn(index, b.array[cursor])
			index++
		}
		for cursor := 0; cursor < b.tail; cursor++ {
			mapfn(index, b.array[cursor])
			index++
		}
	}
}

// String returns a string representation for value buffers.
func (b *Buffer) String() string {
	var values []string
	for _, elem := range b.Array() {
		values = append(values, fmt.Sprintf("%v", elem))
	}
	return strings.Join(values, " <= ")
}

// --------------------------------------------------------------------------------
// Util methods
// --------------------------------------------------------------------------------

func arrayClear(source []float64, index, length int) {
	for x := 0; x < length; x++ {
		absoluteIndex := x + index
		source[absoluteIndex] = 0
	}
}

func arrayCopy(source []float64, sourceIndex int, destination []float64, destinationIndex, length int) {
	for x := 0; x < length; x++ {
		from := sourceIndex + x
		to := destinationIndex + x

		destination[to] = source[from]
	}
}
