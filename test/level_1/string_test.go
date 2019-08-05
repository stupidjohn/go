package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_decode_string(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteString("hello")
		iter := c.CreateIterator(buf.Bytes())
		should.Equal("hello", iter.ReadString())
	}
}

func Test_unmarshal_string(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteString("hello")
		var val string
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal("hello", val)
	}
}

func Test_encode_string(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		stream := c.CreateStream()
		stream.WriteString("hello")
		iter := c.CreateIterator(stream.Buffer())
		should.Equal("hello", iter.ReadString())
	}
}

func Test_marshal_string(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal("hello")
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal("hello", iter.ReadString())
	}
}
