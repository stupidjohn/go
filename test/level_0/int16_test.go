package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_decode_int16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI16(-1)
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(int16(-1), iter.ReadInt16())
	}
}

func Test_unmarshal_int16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI16(-1)
		var val int16
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(int16(-1), val)
	}
}

func Test_encode_int16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteInt16(-1)
		iter := c.CreateIterator(stream.Buffer())
		should.Equal(int16(-1), iter.ReadInt16())
	}
}

func Test_marshal_int16(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal(int16(-1))
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(int16(-1), iter.ReadInt16())
	}
}
