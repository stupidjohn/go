package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go/test"
)

func Test_skip_map_of_string_key(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.STRING, thrift.I64, 1)
		proto.WriteString("1")
		proto.WriteI64(1)
		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipMap(nil))
	}
}

func Test_skip_map_of_string_elem(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.I64, thrift.STRING, 1)
		proto.WriteI64(1)
		proto.WriteString("1")
		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipMap(nil))
	}
}

func Test_decode_map_of_string_key(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteMapBegin(thrift.STRING, thrift.I64, 1)
		proto.WriteString("1")
		proto.WriteI64(1)
		proto.WriteMapEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(map[interface{}]interface{}{
			"1": int64(1),
		}, iter.ReadMap())
	}
}

func Test_encode_map_of_string_key(t *testing.T) {
	should := require.New(t)
	stream := thrifter.NewBufferedStream(nil)
	stream.WriteMap(map[interface{}]interface{}{
		"1": int64(1),
	})
	iter := thrifter.NewBufferedIterator(stream.Buffer())
	should.Equal(map[interface{}]interface{}{
		"1": int64(1),
	}, iter.ReadMap())
}
