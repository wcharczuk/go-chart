package matrix

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestNew(t *testing.T) {
	// replaced new assertions helper

	m := New(10, 5)
	rows, cols := m.Size()
	testutil.AssertEqual(t, 10, rows)
	testutil.AssertEqual(t, 5, cols)
	testutil.AssertZero(t, m.Get(0, 0))
	testutil.AssertZero(t, m.Get(9, 4))
}

func TestNewWithValues(t *testing.T) {
	// replaced new assertions helper

	m := New(5, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	rows, cols := m.Size()
	testutil.AssertEqual(t, 5, rows)
	testutil.AssertEqual(t, 2, cols)
	testutil.AssertEqual(t, 1, m.Get(0, 0))
	testutil.AssertEqual(t, 10, m.Get(4, 1))
}

func TestIdentitiy(t *testing.T) {
	// replaced new assertions helper

	id := Identity(5)
	rows, cols := id.Size()
	testutil.AssertEqual(t, 5, rows)
	testutil.AssertEqual(t, 5, cols)
	testutil.AssertEqual(t, 1, id.Get(0, 0))
	testutil.AssertEqual(t, 1, id.Get(1, 1))
	testutil.AssertEqual(t, 1, id.Get(2, 2))
	testutil.AssertEqual(t, 1, id.Get(3, 3))
	testutil.AssertEqual(t, 1, id.Get(4, 4))
	testutil.AssertEqual(t, 0, id.Get(0, 1))
	testutil.AssertEqual(t, 0, id.Get(1, 0))
	testutil.AssertEqual(t, 0, id.Get(4, 0))
	testutil.AssertEqual(t, 0, id.Get(0, 4))
}

func TestNewFromArrays(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})
	testutil.AssertNotNil(t, m)

	rows, cols := m.Size()
	testutil.AssertEqual(t, 2, rows)
	testutil.AssertEqual(t, 4, cols)
}

func TestOnes(t *testing.T) {
	// replaced new assertions helper

	ones := Ones(5, 10)
	rows, cols := ones.Size()
	testutil.AssertEqual(t, 5, rows)
	testutil.AssertEqual(t, 10, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			testutil.AssertEqual(t, 1, ones.Get(row, col))
		}
	}
}

func TestMatrixEpsilon(t *testing.T) {
	// replaced new assertions helper

	ones := Ones(2, 2)
	ones = ones.WithEpsilon(0.001)
	testutil.AssertEqual(t, 0.001, ones.Epsilon())
}

func TestMatrixArrays(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	testutil.AssertNotNil(t, m)

	arrays := m.Arrays()

	testutil.AssertEqual(t, arrays, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
}

func TestMatrixIsSquare(t *testing.T) {
	// replaced new assertions helper

	testutil.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}).IsSquare())

	testutil.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}).IsSquare())

	testutil.AssertTrue(t, NewFromArrays([][]float64{
		{1, 2},
		{3, 4},
	}).IsSquare())
}

func TestMatrixIsSymmetric(t *testing.T) {
	// replaced new assertions helper

	testutil.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{2, 1, 2},
	}).IsSymmetric())

	testutil.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}).IsSymmetric())

	testutil.AssertTrue(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{2, 1, 2},
		{3, 2, 1},
	}).IsSymmetric())

}

func TestMatrixGet(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	testutil.AssertEqual(t, 1, m.Get(0, 0))
	testutil.AssertEqual(t, 2, m.Get(0, 1))
	testutil.AssertEqual(t, 3, m.Get(0, 2))
	testutil.AssertEqual(t, 4, m.Get(1, 0))
	testutil.AssertEqual(t, 5, m.Get(1, 1))
	testutil.AssertEqual(t, 6, m.Get(1, 2))
	testutil.AssertEqual(t, 7, m.Get(2, 0))
	testutil.AssertEqual(t, 8, m.Get(2, 1))
	testutil.AssertEqual(t, 9, m.Get(2, 2))
}

func TestMatrixSet(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m.Set(1, 1, 99)
	testutil.AssertEqual(t, 99, m.Get(1, 1))
}

func TestMatrixCol(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	testutil.AssertEqual(t, []float64{1, 4, 7}, m.Col(0))
	testutil.AssertEqual(t, []float64{2, 5, 8}, m.Col(1))
	testutil.AssertEqual(t, []float64{3, 6, 9}, m.Col(2))
}

func TestMatrixRow(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	testutil.AssertEqual(t, []float64{1, 2, 3}, m.Row(0))
	testutil.AssertEqual(t, []float64{4, 5, 6}, m.Row(1))
	testutil.AssertEqual(t, []float64{7, 8, 9}, m.Row(2))
}

func TestMatrixSwapRows(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m.SwapRows(0, 1)

	testutil.AssertEqual(t, []float64{4, 5, 6}, m.Row(0))
	testutil.AssertEqual(t, []float64{1, 2, 3}, m.Row(1))
	testutil.AssertEqual(t, []float64{7, 8, 9}, m.Row(2))
}

func TestMatrixCopy(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m2 := m.Copy()
	testutil.AssertFalse(t, m == m2)
	testutil.AssertTrue(t, m.Equals(m2))
}

func TestMatrixDiagonalVector(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	diag := m.DiagonalVector()
	testutil.AssertEqual(t, []float64{1, 2, 3}, diag)
}

func TestMatrixDiagonalVectorLandscape(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7, 99},
		{4, 2, 8, 99},
	})

	diag := m.DiagonalVector()
	testutil.AssertEqual(t, []float64{1, 2}, diag)
}

func TestMatrixDiagonalVectorPortrait(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4},
		{4, 2},
		{99, 99},
	})

	diag := m.DiagonalVector()
	testutil.AssertEqual(t, []float64{1, 2}, diag)
}

func TestMatrixDiagonal(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	m2 := NewFromArrays([][]float64{
		{1, 0, 0},
		{0, 2, 0},
		{0, 0, 3},
	})

	testutil.AssertTrue(t, m.Diagonal().Equals(m2))
}

func TestMatrixEquals(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	testutil.AssertFalse(t, m.Equals(nil))
	var nilMatrix *Matrix
	testutil.AssertTrue(t, nilMatrix.Equals(nil))
	testutil.AssertFalse(t, m.Equals(New(1, 1)))
	testutil.AssertFalse(t, m.Equals(New(3, 3)))
	testutil.AssertTrue(t, m.Equals(New(3, 3, 1, 4, 7, 4, 2, 8, 7, 8, 3)))
}

func TestMatrixL(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	l := m.L()
	testutil.AssertTrue(t, l.Equals(New(3, 3, 1, 2, 3, 0, 5, 6, 0, 0, 9)))
}

func TestMatrixU(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	u := m.U()
	testutil.AssertTrue(t, u.Equals(New(3, 3, 0, 0, 0, 4, 0, 0, 7, 8, 0)))
}

func TestMatrixString(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	testutil.AssertEqual(t, "1 2 3 \n4 5 6 \n7 8 9 \n", m.String())
}

func TestMatrixLU(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 3, 5},
		{2, 4, 7},
		{1, 1, 0},
	})

	l, u, p := m.LU()
	testutil.AssertNotNil(t, l)
	testutil.AssertNotNil(t, u)
	testutil.AssertNotNil(t, p)
}

func TestMatrixQR(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{12, -51, 4},
		{6, 167, -68},
		{-4, 24, -41},
	})

	q, r := m.QR()
	testutil.AssertNotNil(t, q)
	testutil.AssertNotNil(t, r)
}

func TestMatrixTranspose(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	m2 := m.Transpose()

	rows, cols := m2.Size()
	testutil.AssertEqual(t, 3, rows)
	testutil.AssertEqual(t, 4, cols)

	testutil.AssertEqual(t, 1, m2.Get(0, 0))
	testutil.AssertEqual(t, 10, m2.Get(0, 3))
	testutil.AssertEqual(t, 3, m2.Get(2, 0))
}
