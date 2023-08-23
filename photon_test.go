package photon

import (
	"bytes"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type msg struct {
	Field uint64
}

type msg1 struct {
	Field1 uint64
	Field2 msg2
	Field3 [9]uint8
}

type msg2 struct {
	Field1 bool
	Field2 uint32
}

func TestValue(t *testing.T) {
	msg := &msg1{
		Field1: 7,
		Field2: msg2{
			Field1: true,
			Field2: 255,
		},
		Field3: [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	requireT := require.New(t)

	photon1 := NewFromValue(msg)
	requireT.Equal(msg, photon1.V)
	requireT.Equal(
		[]byte{0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		photon1.B,
	)

	photon2 := NewFromBytes[msg1](photon1.B)
	requireT.Equal(photon1.B, photon2.B)
	requireT.Equal(msg, photon2.V)

	buf := bytes.NewBuffer(photon1.B)
	photon3, err := NewFromReader[msg1](buf)
	requireT.NoError(err)
	requireT.Equal(photon1.B, photon3.B)
	requireT.Equal(msg, photon3.V)
}

func TestFromTooSmallBytes(t *testing.T) {
	requireT := require.New(t)

	requireT.Panics(func() {
		NewFromBytes[msg](make([]byte, 1))
	})
}

func TestFromNilBytes(t *testing.T) {
	type msg struct {
		Field uint64
	}

	requireT := require.New(t)

	requireT.Panics(func() {
		NewFromBytes[msg](nil)
	})
}

func TestFromReaderError(t *testing.T) {
	type msg struct {
		Field uint64
	}
	buf := bytes.NewBuffer(make([]byte, 1))

	requireT := require.New(t)

	ph, err := NewFromReader[msg](buf)
	requireT.ErrorIs(err, io.EOF)
	requireT.Nil(ph.B)
	requireT.Nil(ph.V)
}

func TestSlice(t *testing.T) {
	msgs := []msg1{
		{
			Field1: 7,
			Field2: msg2{
				Field1: true,
				Field2: 255,
			},
			Field3: [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			Field1: 2,
			Field2: msg2{
				Field1: false,
				Field2: 100,
			},
			Field3: [9]uint8{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			Field1: 4,
			Field2: msg2{
				Field1: true,
				Field2: 50,
			},
			Field3: [9]uint8{10, 11, 12, 13, 14, 15, 16, 17, 18},
		},
	}

	requireT := require.New(t)

	photon1 := NewFromSlice(msgs)
	requireT.Equal(msgs, photon1.V)
	requireT.Equal(
		[]byte{
			0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xff, 0x0, 0x0, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0x9, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x32, 0x0, 0x0, 0x0, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10, 0x11, 0x12, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		},
		photon1.B,
	)

	photon2 := NewSliceFromBytes[msg1](photon1.B)
	requireT.Equal(photon1.B, photon2.B)
	requireT.Equal(msgs, photon2.V)

	buf := bytes.NewBuffer(photon1.B)
	photon3, err := NewSliceFromReader[msg1](buf, len(msgs))
	requireT.NoError(err)
	requireT.Equal(photon1.B, photon3.B)
	requireT.Equal(msgs, photon3.V)
}

func TestSliceFromTooSmallBytes(t *testing.T) {
	requireT := require.New(t)

	requireT.Panics(func() {
		NewSliceFromBytes[msg](make([]byte, 1))
	})
}

func TestSliceFromNilBytes(t *testing.T) {
	requireT := require.New(t)

	requireT.Panics(func() {
		NewSliceFromBytes[msg](nil)
	})
}

func TestFromNilSlice(t *testing.T) {
	requireT := require.New(t)

	requireT.Panics(func() {
		NewFromSlice[msg](nil)
	})
}

func TestFromEmptySlice(t *testing.T) {
	requireT := require.New(t)

	requireT.Panics(func() {
		NewFromSlice[msg]([]msg{})
	})
}

func TestBytesAreUpdated(t *testing.T) {
	requireT := require.New(t)

	photon1 := NewFromValue(&msg{
		Field: 1,
	})

	byteCopy := make([]byte, len(photon1.B))
	copy(byteCopy, photon1.B)

	photon1.V.Field++
	requireT.NotEqual(byteCopy, photon1.B)

	copy(byteCopy, photon1.B)
	photon2 := NewFromBytes[msg](photon1.B)
	photon2.V.Field++
	requireT.NotEqual(byteCopy, photon2.B)

	buf := bytes.NewBuffer(photon2.B)
	photon3, err := NewFromReader[msg](buf)
	requireT.NoError(err)
	photon3.V.Field++
	requireT.NotEqual(photon2.B, photon3.B)
}

func TestSliceBytesAreUpdated(t *testing.T) {
	requireT := require.New(t)

	photon1 := NewFromSlice([]msg{
		{
			Field: 1,
		},
	})

	byteCopy := make([]byte, len(photon1.B))
	copy(byteCopy, photon1.B)

	photon1.V[0].Field++
	requireT.NotEqual(byteCopy, photon1.B)

	copy(byteCopy, photon1.B)
	photon2 := NewSliceFromBytes[msg](photon1.B)
	photon2.V[0].Field++
	requireT.NotEqual(byteCopy, photon2.B)

	buf := bytes.NewBuffer(photon2.B)
	photon3, err := NewSliceFromReader[msg](buf, 1)
	requireT.NoError(err)
	photon3.V[0].Field++
	requireT.NotEqual(photon2.B, photon3.B)
}

type errReader struct {
}

func (r errReader) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
}

func TestReaderError(t *testing.T) {
	requireT := require.New(t)

	_, err := NewFromReader[msg](errReader{})
	requireT.Error(err)

	_, err = NewSliceFromReader[msg](errReader{}, 1)
	requireT.Error(err)
}

type chunkReader struct {
	buf *bytes.Buffer
}

func (r chunkReader) Read(p []byte) (int, error) {
	return r.buf.Read(p[:1])
}

func TestChunkedReader(t *testing.T) {
	requireT := require.New(t)

	photon1 := NewFromValue(&msg{
		Field: 1,
	})

	r := chunkReader{
		buf: bytes.NewBuffer(photon1.B),
	}

	photon2, err := NewFromReader[msg](r)
	requireT.NoError(err)
	requireT.Equal(photon1.B, photon2.B)
}

func TestSliceChunkedReader(t *testing.T) {
	requireT := require.New(t)

	photon1 := NewFromSlice([]msg{
		{
			Field: 1,
		},
	})

	r := chunkReader{
		buf: bytes.NewBuffer(photon1.B),
	}

	photon2, err := NewSliceFromReader[msg](r, 1)
	requireT.NoError(err)
	requireT.Equal(photon1.B, photon2.B)
}
