//nolint:ineffassign,staticcheck,wastedassign
package photon

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"testing"
	"unsafe"
)

// go test -benchtime=1000x -bench=. -run=^$ -cpuprofile profile.out
// go tool pprof -http="localhost:8000" pprofbin ./profile.out

func BenchmarkNewFromValue(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}

	var v Union[*msg]

	b.StartTimer()
	for range b.N {
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
		v = NewFromValue(m)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkNewFromBytes(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	raw := ph.B

	var v Union[*msg]

	b.StartTimer()
	for range b.N {
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
		v = NewFromBytes[msg](raw)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkNewFromReader(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, 10*b.N))

	var v Union[*msg]

	b.StartTimer()
	for range b.N {
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
		v, _ = NewFromReader[msg](buf)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkNewFromSlice(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}

	var v Union[[]msg]

	b.StartTimer()
	for range b.N {
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
		v = NewFromSlice(m)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkNewSliceFromBytes(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	raw := ph.B

	var v Union[[]msg]

	b.StartTimer()
	for range b.N {
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
		v = NewSliceFromBytes[msg](raw, len(ph.V))
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkNewSliceFromReader(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	buf := bytes.NewBuffer(bytes.Repeat(ph.B, 10*b.N))

	var v Union[[]msg]

	b.StartTimer()
	for range b.N {
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
		v, _ = NewSliceFromReader[msg](buf, 1)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
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

	var v *msg

	b.StartTimer()
	for range b.N {
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
		v = FromBytes[msg](raw)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
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

	var v []msg

	b.StartTimer()
	for range b.N {
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
		v = SliceFromBytes[msg](raw, len(ph.V))
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkFromPointer(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := &msg{Field: math.MaxUint64}
	ph := NewFromValue(m)
	raw := unsafe.Pointer(&ph.B[0])

	var v *msg

	b.StartTimer()
	for range b.N {
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
		v = FromPointer[msg](raw)
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}

func BenchmarkSliceFromPointer(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	type msg struct {
		Field uint64
	}

	m := []msg{{Field: math.MaxUint64}}
	ph := NewFromSlice(m)
	raw := unsafe.Pointer(&ph.B[0])

	var v []msg

	b.StartTimer()
	for range b.N {
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
		v = SliceFromPointer[msg](raw, len(ph.V))
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, v)
}
