package matrix

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

const (
	// DefaultEpsilon represents the minimum precision for matrix math operations.
	DefaultEpsilon = 0.000001
)

var (
	// ErrDimensionMismatch is a typical error.
	ErrDimensionMismatch = errors.New("matrix is not square, cannot invert")
)

// New returns a new matrix.
func New(rows, cols int) *Matrix {
	return &Matrix{
		rows:     rows,
		cols:     cols,
		epsilon:  DefaultEpsilon,
		elements: make([]float64, rows*cols),
	}
}

// Identity returns the identity matrix of a given order.
func Identity(order int) *Matrix {
	m := New(order, order)
	for i := 0; i < order; i++ {
		m.Set(i, i, 1)
	}
	return m
}

// Zeros returns a matrix of a given size zeroed.
func Zeros(rows, cols int) *Matrix {
	return New(rows, cols)
}

// Ones returns an matrix of ones.
func Ones(rows, cols int) *Matrix {
	ones := make([]float64, rows*cols)
	for i := 0; i < (rows * cols); i++ {
		ones[i] = 1
	}

	return &Matrix{
		rows:     rows,
		cols:     cols,
		epsilon:  DefaultEpsilon,
		elements: ones,
	}
}

// NewFromArrays creates a matrix from a jagged array set.
func NewFromArrays(a [][]float64) *Matrix {
	rows := len(a)
	if rows == 0 {
		return nil
	}
	cols := len(a[0])
	m := New(rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			m.Set(row, col, a[row][col])
		}
	}
	return m
}

// Matrix represents a 2d dense array of floats.
type Matrix struct {
	epsilon    float64
	elements   []float64
	rows, cols int
}

// Epsilon returns the maximum precision for math operations.
func (m *Matrix) Epsilon() float64 {
	return m.epsilon
}

// WithEpsilon sets the epsilon on the matrix and returns a reference to the matrix.
func (m *Matrix) WithEpsilon(epsilon float64) *Matrix {
	m.epsilon = epsilon
	return m
}

// Arrays returns the matrix as a two dimensional jagged array.
func (m *Matrix) Arrays() [][]float64 {
	a := make([][]float64, m.rows, m.cols)

	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			a[row][col] = m.Get(row, col)
		}
	}
	return a
}

// Size returns the dimensions of the matrix.
func (m *Matrix) Size() (rows, cols int) {
	rows = m.rows
	cols = m.cols
	return
}

// IsSquare returns if the row count is equal to the column count.
func (m *Matrix) IsSquare() bool {
	return m.rows == m.cols
}

// IsSymmetric returns if the matrix is symmetric about its diagonal.
func (m *Matrix) IsSymmetric() bool {
	if m.rows != m.cols {
		return false
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < i; j++ {
			if m.Get(i, j) != m.Get(j, i) {
				return false
			}
		}
	}
	return true
}

// Get returns the element at the given row, col.
func (m *Matrix) Get(row, col int) float64 {
	index := (m.cols * row) + col
	return m.elements[index]
}

// Set sets a value.
func (m *Matrix) Set(row, col int, val float64) {
	index := (m.cols * row) + col
	m.elements[index] = val
}

// Col returns a column of the matrix as a vector.
func (m *Matrix) Col(col int) Vector {
	values := make([]float64, m.rows)
	for row := 0; row < m.rows; row++ {
		values[col] = m.Get(row, col)
	}
	return Vector(values)
}

// Row returns a row of the matrix as a vector.
func (m *Matrix) Row(row int) Vector {
	values := make([]float64, m.cols)
	for col := 0; col < m.cols; col++ {
		values[col] = m.Get(row, col)
	}
	return Vector(values)
}

// Copy returns a duplicate of a given matrix.
func (m *Matrix) Copy() *Matrix {
	m2 := New(m.rows, m.cols)
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			m2.Set(row, col, m.Get(row, col))
		}
	}
	return m2
}

// DiagonalVector returns a vector from the diagonal of a matrix.
func (m *Matrix) DiagonalVector() Vector {
	rank := minInt(m.rows, m.cols)
	values := make([]float64, rank)

	for index := 0; index < rank; index++ {
		values[index] = m.Get(index, index)
	}
	return Vector(values)
}

// Equals returns if a matrix equals another matrix.
func (m *Matrix) Equals(other *Matrix) bool {
	if other == nil && m != nil {
		return false
	} else if other == nil {
		return true
	}

	if otherRows, otherCols := other.Size(); otherRows != m.rows || otherCols != m.cols {
		return false
	}

	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			if m.Get(row, col) != other.Get(row, col) {
				return false
			}

		}
	}
	return true
}

// L returns the matrix with zeros below the diagonal.
func (m *Matrix) L() *Matrix {
	m2 := New(m.rows, m.cols)
	for row := 0; row < m.rows; row++ {
		for col := row; col < m.cols; col++ {
			m2.Set(row, col, m.Get(row, col))
		}
	}
	return m2
}

// U returns the matrix with zeros above the diagonal.
func (m *Matrix) U() *Matrix {
	m2 := New(m.rows, m.cols)
	for row := 0; row < m.rows; row++ {
		for col := 0; col < row && col < m.cols; col++ {
			m2.Set(row, col, m.Get(row, col))
		}
	}
	return m2
}

// Diagonal returns a matrix from the diagonal of a matrix.
func (m *Matrix) Diagonal() *Matrix {
	rank := minInt(m.rows, m.cols)
	m2 := New(rank, rank)

	for index := 0; index < rank; index++ {
		m2.Set(index, index, m.Get(index, index))
	}
	return m2
}

// String returns a string representation of the matrix.
func (m *Matrix) String() string {
	buffer := bytes.NewBuffer(nil)
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			buffer.WriteString(fmt.Sprintf("%f", m.Get(row, col)))
			buffer.WriteRune(' ')
		}
		buffer.WriteRune('\n')
	}
	return buffer.String()
}

// Decompositions

// LU returns the LU decomposition of a matrix.
func (m *Matrix) LU() (l, u, p *Matrix) {
	return
}

// QR performs the qr decomposition.
func (m *Matrix) QR() (q, r *Matrix) {
	rows, cols := m.Size()
	qr := m.Copy()
	q = New(rows, cols)
	r = New(rows, cols)

	var i, j, k int
	var norm, s float64

	for k = 0; k < cols; k++ {
		norm = 0
		for i = k; i < rows; i++ {
			norm = math.Hypot(norm, qr.Get(i, k))
		}

		if norm != 0 {
			if qr.Get(k, k) < 0 {
				norm = -norm
			}

			for i = k; i < rows; i++ {
				qr.Set(i, k, qr.Get(i, k)/norm)
			}
			qr.Set(k, k, qr.Get(k, k)+1.0)

			for j = k + 1; j < cols; j++ {
				s = 0
				for i = k; i < rows; i++ {
					s += qr.Get(i, k) * qr.Get(i, j)
				}
				s = -s / qr.Get(k, k)
				for i = k; i < rows; i++ {
					qr.Set(i, j, qr.Get(i, j)+s*qr.Get(i, k))

					if i < j {
						r.Set(i, j, qr.Get(i, j))
					}
				}

			}
		}

		r.Set(k, k, -norm)

	}

	//Q Matrix:
	i, j, k = 0, 0, 0

	for k = cols - 1; k >= 0; k-- {
		q.Set(k, k, 1.0)
		for j = k; j < cols; j++ {
			if qr.Get(k, k) != 0 {
				s = 0
				for i = k; i < rows; i++ {
					s += qr.Get(i, k) * q.Get(i, j)
				}
				s = -s / qr.Get(k, k)
				for i = k; i < rows; i++ {
					q.Set(i, j, q.Get(i, j)+s*qr.Get(i, k))
				}
			}
		}
	}

	return
}
