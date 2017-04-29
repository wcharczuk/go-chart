package util

import (
	"fmt"
	"strings"
)

const (
	valueBufferMinimumGrow     = 4
	valueBufferShrinkThreshold = 32
	valueBufferGrowFactor      = 200
	valueBufferDefaultCapacity = 4
)

var (
	emptyArray = make([]float64, 0)
)

// NewValueBuffer creates a new value buffer with an optional set of values.
func NewValueBuffer(values ...float64) *ValueBuffer {
	var tail int
	array := make([]float64, Math.MaxInt(len(values), valueBufferDefaultCapacity))
	if len(values) > 0 {
		copy(array, values)
		tail = len(values)
	}
	return &ValueBuffer{
		array: array,
		head:  0,
		tail:  tail,
		size:  len(values),
	}
}

// NewValueBufferWithCapacity creates a new ValueBuffer pre-allocated with the given capacity.
func NewValueBufferWithCapacity(capacity int) *ValueBuffer {
	return &ValueBuffer{
		array: make([]float64, capacity),
		head:  0,
		tail:  0,
		size:  0,
	}
}

// ValueBuffer is a fifo buffer that is backed by a pre-allocated array, instead of allocating
// a whole new node object for each element (which saves GC churn).
// Enqueue can be O(n), Dequeue can be O(1).
type ValueBuffer struct {
	array []float64
	head  int
	tail  int
	size  int
}

// Len returns the length of the ValueBuffer (as it is currently populated).
// Actual memory footprint may be different.
func (vb *ValueBuffer) Len() int {
	return vb.size
}

// GetValue implements sequence provider.
func (vb *ValueBuffer) GetValue(index int) float64 {
	effectiveIndex := (vb.head + index) % len(vb.array)
	return vb.array[effectiveIndex]
}

// Capacity returns the total size of the ValueBuffer, including empty elements.
func (vb *ValueBuffer) Capacity() int {
	return len(vb.array)
}

// SetCapacity sets the capacity of the ValueBuffer.
func (vb *ValueBuffer) SetCapacity(capacity int) {
	newArray := make([]float64, capacity)
	if vb.size > 0 {
		if vb.head < vb.tail {
			arrayCopy(vb.array, vb.head, newArray, 0, vb.size)
		} else {
			arrayCopy(vb.array, vb.head, newArray, 0, len(vb.array)-vb.head)
			arrayCopy(vb.array, 0, newArray, len(vb.array)-vb.head, vb.tail)
		}
	}
	vb.array = newArray
	vb.head = 0
	if vb.size == capacity {
		vb.tail = 0
	} else {
		vb.tail = vb.size
	}
}

// Clear removes all objects from the ValueBuffer.
func (vb *ValueBuffer) Clear() {
	if vb.head < vb.tail {
		arrayClear(vb.array, vb.head, vb.size)
	} else {
		arrayClear(vb.array, vb.head, len(vb.array)-vb.head)
		arrayClear(vb.array, 0, vb.tail)
	}

	vb.head = 0
	vb.tail = 0
	vb.size = 0
}

// Enqueue adds an element to the "back" of the ValueBuffer.
func (vb *ValueBuffer) Enqueue(value float64) {
	if vb.size == len(vb.array) {
		newCapacity := int(len(vb.array) * int(valueBufferGrowFactor/100))
		if newCapacity < (len(vb.array) + valueBufferMinimumGrow) {
			newCapacity = len(vb.array) + valueBufferMinimumGrow
		}
		vb.SetCapacity(newCapacity)
	}

	vb.array[vb.tail] = value
	vb.tail = (vb.tail + 1) % len(vb.array)
	vb.size++
}

// Dequeue removes the first element from the RingBuffer.
func (vb *ValueBuffer) Dequeue() float64 {
	if vb.size == 0 {
		return 0
	}

	removed := vb.array[vb.head]
	vb.head = (vb.head + 1) % len(vb.array)
	vb.size--
	return removed
}

// Peek returns but does not remove the first element.
func (vb *ValueBuffer) Peek() float64 {
	if vb.size == 0 {
		return 0
	}
	return vb.array[vb.head]
}

// PeekBack returns but does not remove the last element.
func (vb *ValueBuffer) PeekBack() float64 {
	if vb.size == 0 {
		return 0
	}
	if vb.tail == 0 {
		return vb.array[len(vb.array)-1]
	}
	return vb.array[vb.tail-1]
}

// TrimExcess resizes the buffer to better fit the contents.
func (vb *ValueBuffer) TrimExcess() {
	threshold := float64(len(vb.array)) * 0.9
	if vb.size < int(threshold) {
		vb.SetCapacity(vb.size)
	}
}

// Array returns the ring buffer, in order, as an array.
func (vb *ValueBuffer) Array() Array {
	newArray := make([]float64, vb.size)

	if vb.size == 0 {
		return newArray
	}

	if vb.head < vb.tail {
		arrayCopy(vb.array, vb.head, newArray, 0, vb.size)
	} else {
		arrayCopy(vb.array, vb.head, newArray, 0, len(vb.array)-vb.head)
		arrayCopy(vb.array, 0, newArray, len(vb.array)-vb.head, vb.tail)
	}

	return Array(newArray)
}

// Each calls the consumer for each element in the buffer.
func (vb *ValueBuffer) Each(mapfn func(int, float64)) {
	if vb.size == 0 {
		return
	}

	var index int
	if vb.head < vb.tail {
		for cursor := vb.head; cursor < vb.tail; cursor++ {
			mapfn(index, vb.array[cursor])
			index++
		}
	} else {
		for cursor := vb.head; cursor < len(vb.array); cursor++ {
			mapfn(index, vb.array[cursor])
			index++
		}
		for cursor := 0; cursor < vb.tail; cursor++ {
			mapfn(index, vb.array[cursor])
			index++
		}
	}
}

// String returns a string representation for value buffers.
func (vb *ValueBuffer) String() string {
	var values []string
	for _, elem := range vb.Array() {
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
