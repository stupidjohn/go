package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_decode_uint16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI16(1024)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(uint16(1024), iter.ReadUint16())
	}
}

func Test_unmarshal_uint16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI16(1024)
		var val uint16
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(uint16(1024), val)
	}
}

func Test_encode_uint16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteUint16(1024)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(uint16(1024), iter.ReadUint16())
	}
}

func Test_marshal_uint16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(uint16(1024))
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(uint16(1024), iter.ReadUint16())
	}
}
