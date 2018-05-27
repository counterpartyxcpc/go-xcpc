package playground

import (
	"encoding/hex"
	"fmt"
	enc "go-xcpc/message"
	xm "go-xcpc/proto"

	proto "github.com/golang/protobuf/proto"
)

var key = "1789916b3326647d973606cdbd15c4aa127e4b7079bccf599905e14a9f22593a"
var prefix = []byte{0x00, 0x58, 0x43, 0x50} //XCP prefix

// Experiment test
func Experiment() {
	var enctx []byte
	var hexkey []byte

	prototx := xm.XCPCsend{Asset: "XCPC", Quantity: 100, Address: "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9", Memo: "this is the test"}
	// plaintx is the plain protobuf message
	plaintx, err := proto.Marshal(&prototx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plain protobuf sample msg:\n%s\n", plaintx)

	hexkey, err = hex.DecodeString(key)
	if err != nil {
		panic(err)
	}

	encmsg, err := enc.Rc4Enc(hexkey, plaintx)
	if err != nil {
		panic(err)
	}
	enctx = append(enctx, prefix...)
	enctx = append(enctx, encmsg...)
	fmt.Printf("RC4 encrypted msg:\n%s\n.\n", enctx)
}

// ProtoOneof test the oneof module of protobuf
func ProtoOneof() {
	var xcpcsend = &xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Send{
			&xm.XCPCsend{Asset: "XCPC", Quantity: 100, Address: "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9", Memo: "this is test"},
		},
	}

	var xcpcbroad = &xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Broadcast{
			&xm.XCPCbroadcast{Text: "XCPC", Value: 100, Feefraction: 1000, Timestamp: "27-May-2018"},
		},
	}
	fmt.Println(xcpcsend.Msgtype)
	fmt.Println(xcpcbroad.Msgtype)
}
