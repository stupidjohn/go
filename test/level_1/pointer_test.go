package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stupidjohn/go/test"
	"testing"
)

func Test_unmarshal_ptr_int64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.UnmarshalCombinations {
		buf, proto := c.CreateProtocol()
		proto.WriteI64(2)
		proto.WriteListEnd()
		var val *int64
		should.NoError(c.Unmarshal(buf.Bytes(), &val))
		should.Equal(int64(2), *val)
	}
}
func Test_marshal_ptr_int64(t *testing.T) {
	should := require.New(t)
	for _, c := range test.MarshalCombinations {
		val := int64(2)
		output, err := c.Marshal(&val)
		should.NoError(err)
		iter := c.CreateIterator(output)
		should.Equal(int64(2), iter.ReadInt64())
	}
}
