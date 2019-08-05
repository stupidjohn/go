package reflection

import (
	"github.com/stupidjohn/go/spi"
	"reflect"
	"unsafe"
)

type pointerDecoder struct {
	valType    reflect.Type
	valDecoder internalDecoder
}

func (decoder *pointerDecoder) decode(ptr unsafe.Pointer, iter spi.Iterator) {
	value := reflect.New(decoder.valType).Interface()
	newPtr := (*emptyInterface)(unsafe.Pointer(&value)).word
	decoder.valDecoder.decode(newPtr, iter)
	*(*unsafe.Pointer)(ptr) = newPtr
}
