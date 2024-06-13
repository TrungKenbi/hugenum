package hugenum_test

import (
	"testing"

	"github.com/TrungKenbi/hugenum"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a, b := hugenum.New(10, 0), hugenum.New(1, 1)
	a.Add(b)
	assert.Equal(t, hugenum.New(20, 0), a)

	a, b = hugenum.New(1, 1), hugenum.New(10, 0)
	a.Add(b)
	assert.Equal(t, hugenum.New(20, 0), a)

	a, b = hugenum.New(10, 0), hugenum.New(1, 0)
	a.Add(b)
	assert.Equal(t, hugenum.New(11, 0), a)
}

func TestSub(t *testing.T) {
	a, b := hugenum.New(10, 0), hugenum.New(1, 0)
	a.Subtract(b)
	assert.Equal(t, hugenum.New(9, 0), a)
}

func TestMul(t *testing.T) {
	a := hugenum.New(10, 0)
	a.MultiplyFactor(2)
	assert.Equal(t, hugenum.New(20, 0), a)
}

func TestMulOther(t *testing.T) {
	a := hugenum.New(10, 0)
	b := hugenum.New(2, 0)
	a.Multiply(b)
	assert.Equal(t, hugenum.New(20, 0), a)

	a = hugenum.New(10, 0)
	b = hugenum.New(10, 0)
	a.Multiply(b)
	assert.Equal(t, hugenum.New(100, 0), a)

	a = hugenum.New(1, 1)
	b = hugenum.New(10, 0)
	a.Multiply(b)
	assert.Equal(t, hugenum.New(100, 0), a)

	a = hugenum.New(1, 3)
	b = hugenum.New(1000, 0)
	a.Multiply(b)
	assert.Equal(t, hugenum.New(1, 6), a)
}

func TestDiv(t *testing.T) {
	a := hugenum.New(10, 0)
	a.Divide(2)
	assert.Equal(t, hugenum.New(5, 0), a)

	a = hugenum.New(10, 0)
	a.Divide(3)
	assert.Equal(t, hugenum.New(10/3., 0), a)
}

func TestPowTen(t *testing.T) {
	a := hugenum.New(1000, 2)
	a.PowTen(9)

	assert.Equal(t, hugenum.New(100, 12), a)
}

func TestNegative(t *testing.T) {
	a := hugenum.New(-1000, 2)
	a.PowTen(9)

	assert.Equal(t, hugenum.New(-100, 12), a)
}
