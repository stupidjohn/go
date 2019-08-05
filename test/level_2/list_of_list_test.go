package test

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/general"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_skip_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.LIST, 2)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(2)
		proto.WriteListEnd()
		proto.WriteListEnd()
		iter := c.CreateIterator(buf.Bytes())
		should.Equal(buf.Bytes(), iter.SkipList(nil))
	}
}

func Test_unmarshal_general_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.LIST, 2)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(2)
		proto.WriteListEnd()
		proto.WriteListEnd()
		var val general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_unmarshal_list_of_general_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.LIST, 2)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(2)
		proto.WriteListEnd()
		proto.WriteListEnd()
		var val []general.List
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_unmarshal_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteListBegin(thrift.LIST, 2)
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(1)
		proto.WriteListEnd()
		proto.WriteListBegin(thrift.I64, 1)
		proto.WriteI64(2)
		proto.WriteListEnd()
		proto.WriteListEnd()
		var val [][]int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal([][]int64{
			{1}, {2},
		}, val)
	}
}

func Test_marshal_general_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.Combinations {
		lst := general.List{
			general.List{
				int64(1),
			},
			general.List{
				int64(2),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_marshal_list_of_general_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		lst := []general.List{
			{
				int64(1),
			},
			{
				int64(2),
			},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}

func Test_marshal_list_of_list(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		lst := [][]int64{
			{1}, {2},
		}

		output, err := c.Marshal(lst)
		should.NoError(err)
		output1, err := c.Marshal(&lst)
		should.NoError(err)
		should.Equal(output, output1)
		var val general.List
		should.NoError(c.Unmarshal(output, &val))
		should.Equal(general.List{int64(1)}, val[0])
	}
}
