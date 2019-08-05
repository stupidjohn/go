package reflection

import (
	"github.com/stupidjohn/go/protocol"
	"github.com/stupidjohn/go/spi"
	"reflect"
	"unsafe"
)

type sliceEncoder struct {
	sliceType   reflect.Type
	elemType    reflect.Type
	elemEncoder internalEncoder
}

func (encoder *sliceEncoder) encode(ptr unsafe.Pointer, stream spi.Stream) {
	slice := (*sliceHeader)(ptr)
	stream.WriteListHeader(encoder.elemEncoder.thriftType(), slice.Len)
	offset := uintptr(slice.Data)
	var addr unsafe.Pointer
	for i := 0; i < slice.Len; i++ {
		addr = unsafe.Pointer(offset)
		if encoder.elemType.Kind() == reflect.Map {
			addr = unsafe.Pointer((uintptr)(*(*uint64)(addr)))
		}
		encoder.elemEncoder.encode(addr, stream)
		offset += encoder.elemType.Size()
	}
}

func (encoder *sliceEncoder) thriftType() protocol.TType {
	return protocol.TypeList
}
