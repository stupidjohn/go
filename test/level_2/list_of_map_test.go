package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go/test"
	"github.com/thrift-iterator/go/general"
)

func Test_skip_list_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.MAP, 2)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(1)
		proto.WriteI64(1)
		proto.WriteMapEnd()
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteListEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipList(nil))
	}
}

func Test_unmarshal_general_list_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.MAP, 2)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(1)
		proto.WriteI64(1)
		proto.WriteMapEnd()
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteListEnd()
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.Map{
			int32(1): int64(1),
		}, val[0])
		should.Equal(int64(1), val.Get(0, int32(1)))
	}
}

func Test_unmarshal_list_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.MAP, 2)
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(1)
		proto.WriteI64(1)
		proto.WriteMapEnd()
		proto.WriteMapBegin(thrift.I32, thrift.I64, 1)
		proto.WriteI32(2)
		proto.WriteI64(2)
		proto.WriteMapEnd()
		proto.WriteListEnd()
		var val []map[int32]int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([]map[int32]int64{
			{1: 1}, {2: 2},
		}, val)
	}
}

func Test_marshal_general_list_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		output, err := c.Marshal(general.List{
			general.Map{
				int32(1): int64(1),
			},
			general.Map{
				int32(2): int64(2),
			},
		})
		should.NoError(err)
		var val []map[int32]int64
		should.NoError(c.Unmarshal(output, &val))
		should.Equal([]map[int32]int64{
			{1: 1}, {2: 2},
		}, val)
	}
}

func Test_marshal_list_of_map(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		output, err := c.Marshal([]map[int32]int64{
			{1: 1}, {2: 2},
		})
		should.NoError(err)
		var val []map[int32]int64
		should.NoError(c.Unmarshal(output, &val))
		should.Equal([]map[int32]int64{
			{1: 1}, {2: 2},
		}, val)
	}
}
