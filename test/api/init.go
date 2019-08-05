package test

import (
	"github.com/stupidjohn/go"
	"github.com/stupidjohn/go/test/api/binding_test"
	"github.com/v2pro/wombat/generic"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
}.Froze()

//go:generate go install github.com/stupidjohn/go/cmd/thrifter
//go:generate $GOPATH/bin/thrifter -pkg github.com/stupidjohn/go/test/api
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*binding_test.TestObject)(nil),
		)
	})
}
