package raw

import (
	"github.com/stupidjohn/go/protocol"
	"github.com/stupidjohn/go/spi"
)

type rawMapEncoder struct {
}

func (encoder *rawMapEncoder) Encode(val interface{}, stream spi.Stream) {
	obj := val.(Map)
	length := len(obj.Entries)
	stream.WriteMapHeader(obj.KeyType, obj.ElementType, length)
	for _, entry := range obj.Entries {
		stream.Write(entry.Key)
		stream.Write(entry.Element)
	}
}

func (encoder *rawMapEncoder) ThriftType() protocol.TType {
	return protocol.TypeMap
}
