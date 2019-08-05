package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_decode_uint64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI64(1024)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(uint64(1024), iter.ReadUint64())
	}
}

func Test_unmarshal_uint64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI64(1024)
		var val uint64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(uint64(1024), val)
	}
}

func Test_encode_uint64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteUint64(1024)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(uint64(1024), iter.ReadUint64())
	}
}

func Test_marshal_uint64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(uint64(1024))
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(uint64(1024), iter.ReadUint64())
	}
}
