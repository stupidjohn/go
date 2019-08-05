package test

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/general"
	"github.com/stupidjohn/go/test"
	"testing"
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
		lst := general.List{
			general.Map{
				int32(1): int64(1),
			},
			general.Map{
				int32(2): int64(2),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
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
		lst := []map[int32]int64{
			{1: 1}, {2: 2},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.Equal(output, output1)
		should.NoError(err)
		var val []map[int32]int64
		should.NoError(c.Unmarshal(output, &val))
		should.Equal([]map[int32]int64{
			{1: 1}, {2: 2},
		}, val)
	}
}
