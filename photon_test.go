package photon

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoton(t *testing.T) {
	type msg2 struct {
		Field1 bool
		Field2 uint32
	}

	type msg1 struct {
		Field1 uint64
		Field2 msg2
		Field3 [9]uint8
	}

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
	type msg struct {
		Field uint64
	}

	requireT := require.New(t)

	requireT.Panics(func() {
		NewFromBytes[msg](make([]byte, 1))
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
