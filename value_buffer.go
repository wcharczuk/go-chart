package chart

import (
	"constraints"
	"fmt"
	"strings"
)

const (
	bufferMinimumGrow     = 4
	bufferShrinkThreshold = 32
	bufferGrowFactor      = 200
	bufferDefaultCapacity = 4
)

// NewValueBuffer creates a new value buffer with an optional set of values.
func NewValueBuffer[A constraints.Ordered](values ...A) *ValueBuffer[A] {
	var tail int
	array := make([]A, Max(len(values), bufferDefaultCapacity))
	if len(values) > 0 {
		copy(array, values)
		tail = len(values)
	}
	return &ValueBuffer[A]{
		array: array,
		head:  0,
		tail:  tail,
		size:  len(values),
	}
}

// NewValueBufferWithCapacity creates a new ValueBuffer pre-allocated with the given capacity.
func NewValueBufferWithCapacity[A any](capacity int) *ValueBuffer[A] {
	return &ValueBuffer[A]{
		array: make([]A, capacity),
		head:  0,
		tail:  0,
		size:  0,
	}
}

// ValueBuffer is a fifo data structure that is backed by a pre-allocated array.
// Instead of allocating a whole new node object for each element, array elements are re-used (which saves GC churn).
// Enqueue can be O(n), Dequeue is generally O(1).
// Buffer implements `seq.Provider`
type ValueBuffer[A any] struct {
	array []A
	head  int
	tail  int
	size  int
}

// Len returns the length of the Buffer (as it is currently populated).
// Actual memory footprint may be different.
func (b *ValueBuffer[A]) Len() int {
	return b.size
}

// GetValue implements seq provider.
func (b *ValueBuffer[A]) GetValue(index int) A {
	effectiveIndex := (b.head + index) % len(b.array)
	return b.array[effectiveIndex]
}

// Capacity returns the total size of the Buffer, including empty elements.
func (b *ValueBuffer[A]) Capacity() int {
	return len(b.array)
}

// SetCapacity sets the capacity of the Buffer.
func (b *ValueBuffer[A]) SetCapacity(capacity int) {
	newArray := make([]A, capacity)
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
func (b *ValueBuffer[A]) Clear() {
	b.array = make([]A, bufferDefaultCapacity)
	b.head = 0
	b.tail = 0
	b.size = 0
}

// Enqueue adds an element to the "back" of the Buffer.
func (b *ValueBuffer[A]) Enqueue(value A) {
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
func (b *ValueBuffer[A]) Dequeue() (output A) {
	if b.size == 0 {
		return
	}

	output = b.array[b.head]
	b.head = (b.head + 1) % len(b.array)
	b.size--
	return
}

// Peek returns but does not remove the first element.
func (b *ValueBuffer[A]) Peek() (output A) {
	if b.size == 0 {
		return
	}
	output = b.array[b.head]
	return
}

// PeekBack returns but does not remove the last element.
func (b *ValueBuffer[A]) PeekBack() (output A) {
	if b.size == 0 {
		return
	}
	if b.tail == 0 {
		output = b.array[len(b.array)-1]
		return
	}
	output = b.array[b.tail-1]
	return
}

// TrimExcess resizes the capacity of the buffer to better fit the contents.
func (b *ValueBuffer[A]) TrimExcess() {
	threshold := float64(len(b.array)) * 0.9
	if b.size < int(threshold) {
		b.SetCapacity(b.size)
	}
}

// Array returns the ring buffer, in order, as an array.
func (b *ValueBuffer[A]) Array() []A {
	newArray := make([]A, b.size)

	if b.size == 0 {
		return newArray
	}

	if b.head < b.tail {
		arrayCopy(b.array, b.head, newArray, 0, b.size)
	} else {
		arrayCopy(b.array, b.head, newArray, 0, len(b.array)-b.head)
		arrayCopy(b.array, 0, newArray, len(b.array)-b.head, b.tail)
	}

	return newArray
}

// Each calls the consumer for each element in the buffer.
func (b *ValueBuffer[A]) Each(mapfn func(int, A)) {
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
func (b *ValueBuffer[A]) String() string {
	var values []string
	for _, elem := range b.Array() {
		values = append(values, fmt.Sprint(elem))
	}
	return strings.Join(values, " <= ")
}

// --------------------------------------------------------------------------------
// Util methods
// --------------------------------------------------------------------------------

func arrayClear[A any](source []A, index, length int) {
	var zero A
	for x := 0; x < length; x++ {
		absoluteIndex := x + index
		source[absoluteIndex] = zero
	}
}

func arrayCopy[A any](source []A, sourceIndex int, destination []A, destinationIndex, length int) {
	for x := 0; x < length; x++ {
		from := sourceIndex + x
		to := destinationIndex + x

		destination[to] = source[from]
	}
}
