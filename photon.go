package photon

import (
	"io"
	"unsafe"

	"github.com/pkg/errors"
)

// Union represents a C-like union between bytes and structure.
type Union[T any] struct {
	B []byte
	V T
}

// NewFromBytes creates union from bytes.
func NewFromBytes[T comparable](b []byte) Union[*T] {
	var v T
	return Union[*T]{
		B: b[:unsafe.Sizeof(v)],
		V: (*T)(unsafe.Pointer(&b[0])),
	}
}

// NewFromValue creates union from value.
func NewFromValue[T comparable](v *T) Union[*T] {
	return Union[*T]{
		B: unsafe.Slice((*byte)(unsafe.Pointer(v)), unsafe.Sizeof(*v)),
		V: v,
	}
}

// NewFromReader creates union from bytes read from the reader.
func NewFromReader[T comparable](r io.Reader) (Union[*T], error) {
	var v T
	b := make([]byte, unsafe.Sizeof(v))
	p := Union[*T]{
		B: b,
		V: (*T)(unsafe.Pointer(&b[0])),
	}

	var n int
	var err error
	for {
		n, err = r.Read(b)
		if len(b) == n {
			return p, nil
		}
		if err != nil {
			return Union[*T]{}, errors.WithStack(err)
		}
		b = b[n:]
	}
}

// NewSliceFromBytes creates union containing slice from bytes.
func NewSliceFromBytes[T comparable](b []byte) Union[[]T] {
	var v T
	itemSize := int(unsafe.Sizeof(v))
	n := len(b) / itemSize

	if n == 0 {
		panic("byte slice is too small")
	}

	return Union[[]T]{
		B: b[:itemSize*n],
		V: unsafe.Slice((*T)(unsafe.Pointer(&b[0])), n),
	}
}

// NewFromSlice creates union from slice.
func NewFromSlice[T comparable](v []T) Union[[]T] {
	if len(v) == 0 {
		panic("slice is empty")
	}
	return Union[[]T]{
		B: unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), int(unsafe.Sizeof(v[0]))*len(v)),
		V: v,
	}
}

// NewSliceFromReader creates union containing slice from bytes read from the reader.
func NewSliceFromReader[T comparable](r io.Reader, n int) (Union[[]T], error) {
	var v T

	b := make([]byte, int(unsafe.Sizeof(v))*n)
	p := Union[[]T]{
		B: b,
		V: unsafe.Slice((*T)(unsafe.Pointer(&b[0])), n),
	}

	var err error
	for {
		n, err = r.Read(b)
		if len(b) == n {
			return p, nil
		}
		if err != nil {
			return Union[[]T]{}, errors.WithStack(err)
		}
		b = b[n:]
	}
}
