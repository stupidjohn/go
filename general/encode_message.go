package general

import (
	"github.com/stupidjohn/go/protocol"
	"github.com/stupidjohn/go/spi"
)

type messageEncoder struct {
}

func (encoder *messageEncoder) Encode(val interface{}, stream spi.Stream) {
	msg := val.(Message)
	stream.WriteMessageHeader(msg.MessageHeader)
	writeStruct(msg.Arguments, stream)
}

func (encoder *messageEncoder) ThriftType() protocol.TType {
	return protocol.TypeStruct
}

type messageHeaderEncoder struct {
}

func (encoder *messageHeaderEncoder) Encode(val interface{}, stream spi.Stream) {
	msgHeader := val.(protocol.MessageHeader)
	stream.WriteMessageHeader(msgHeader)
}

func (encoder *messageHeaderEncoder) ThriftType() protocol.TType {
	return protocol.TypeStruct
}
