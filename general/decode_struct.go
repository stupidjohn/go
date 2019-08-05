package general

import (
	"github.com/stupidjohn/go/protocol"
	"github.com/stupidjohn/go/spi"
)

type generalStructDecoder struct {
}

func (decoder *generalStructDecoder) Decode(val interface{}, iter spi.Iterator) {
	*val.(*Struct) = readStruct(iter).(Struct)
}

func readStruct(iter spi.Iterator) interface{} {
	generalStruct := Struct{}
	iter.ReadStructHeader()
	for {
		fieldType, fieldId := iter.ReadStructField()
		if fieldType == protocol.TypeStop {
			return generalStruct
		}
		generalReader := generalReaderOf(fieldType)
		generalStruct[fieldId] = generalReader(iter)
	}
}
