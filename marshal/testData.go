//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf  --gogoslick_out=proto testdata.proto
package marshal

import "github.com/Reshusk23/dme-go-logger/marshal/proto"

// TestData -
type TestData struct {
	proto.TestData
}
