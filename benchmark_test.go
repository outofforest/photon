package photon

import (
	"bytes"
	"math"
	"testing"
)

// go test -bench=. -run=^$ -cpuprofile profile.out
// go tool pprof -http="localhost:8000" pprofbin ./profile.out

func BenchmarkFromValue(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			NewFromValue(m)
		}
		b.StopTimer()
	}
}

func BenchmarkFromBytes(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	raw := ph.B

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			NewFromBytes[msg](raw)
		}
		b.StopTimer()
	}
}

func BenchmarkFromReader(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	const loop = 10000

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, loop))

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < loop; i++ {
			_, _ = NewFromReader[msg](buf)
		}
		b.StopTimer()
	}
}

func BenchmarkFromSlice(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			NewFromSlice(m)
		}
		b.StopTimer()
	}
}

func BenchmarkSliceFromBytes(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	raw := ph.B

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			NewSliceFromBytes[msg](raw)
		}
		b.StopTimer()
	}
}

func BenchmarkSliceFromReader(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	const loop = 10000

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, loop))

	for bi := 0; bi < b.N; bi++ {
		b.StartTimer()
		for i := 0; i < loop; i++ {
			_, _ = NewSliceFromReader[msg](buf, 1)
		}
		b.StopTimer()
	}
}
