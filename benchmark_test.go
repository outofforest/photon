package photon

import (
	"bytes"
	"math"
	"testing"
)

func BenchmarkFromValue(b *testing.B) {
	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		NewFromValue(m)
	}
	b.StopTimer()
}

func BenchmarkFromBytes(b *testing.B) {
	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	raw := ph.B

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		NewFromBytes[msg](raw)
	}
	b.StopTimer()
}

func BenchmarkFromReader(b *testing.B) {
	type msg struct {
		Field uint64
	}

	const loop = 10000

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, loop))

	b.ResetTimer()
	for i := 0; i < loop; i++ {
		_, _ = NewFromReader[msg](buf)
	}
	b.StopTimer()
}

func BenchmarkFromSlice(b *testing.B) {
	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		NewFromSlice(m)
	}
	b.StopTimer()
}

func BenchmarkSliceFromBytes(b *testing.B) {
	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	raw := ph.B

	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		NewSliceFromBytes[msg](raw)
	}
	b.StopTimer()
}

func BenchmarkSliceFromReader(b *testing.B) {
	type msg struct {
		Field uint64
	}

	const loop = 10000

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, loop))

	b.ResetTimer()
	for i := 0; i < loop; i++ {
		_, _ = NewSliceFromReader[msg](buf, 1)
	}
	b.StopTimer()
}
