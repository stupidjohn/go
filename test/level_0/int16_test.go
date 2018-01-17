package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/test"
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

func Test_encode_int16(t *testing.T) {
	should := require.New(t)
	stream := thrifter.NewStream(nil)
	stream.WriteInt16(-1)
	iter := thrifter.NewIterator(nil, stream.Buffer())
	should.Equal(int16(-1), iter.ReadInt16())
}