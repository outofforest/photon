package photon

import (
	"io"
	"unsafe"

	"github.com/pkg/errors"
)

// Union represents a C-like union between bytes and structure.
type Union[T comparable] struct {
	Bytes []byte
	Value *T
}

// NewFromBytes creates union from bytes.
func NewFromBytes[T comparable](b []byte) *Union[T] {
	var v T
	return &Union[T]{
		Bytes: b[:unsafe.Sizeof(v)],
		Value: (*T)(unsafe.Pointer(&b[0])),
	}
}

// NewFromValue creates union from value.
func NewFromValue[T comparable](v *T) *Union[T] {
	return &Union[T]{
		Bytes: unsafe.Slice((*byte)(unsafe.Pointer(v)), unsafe.Sizeof(*v)),
		Value: v,
	}
}

// NewFromReader creates union from bytes read from the reader.
func NewFromReader[T comparable](r io.Reader) (*Union[T], error) {
	var v T
	b := make([]byte, unsafe.Sizeof(v))
	p := &Union[T]{
		Bytes: b,
		Value: (*T)(unsafe.Pointer(&b[0])),
	}

	var n int
	var err error
	for {
		n, err = r.Read(b)
		if len(b) == n {
			return p, nil
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		b = b[n:]
	}
}
