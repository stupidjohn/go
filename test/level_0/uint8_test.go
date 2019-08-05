package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_decode_uint8(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteByte(100)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(uint8(100), iter.ReadUint8())
	}
}

func Test_unmarshal_uint8(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteByte(100)
		var val uint8
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(uint8(100), val)
	}
}

func Test_encode_uint8(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteUint8(100)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(uint8(100), iter.ReadUint8())
	}
}

func Test_marshal_uint8(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(uint8(100))
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(uint8(100), iter.ReadUint8())
	}
}
