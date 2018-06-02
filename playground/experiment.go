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

	prototx := xm.XCPCTransaction_Send{Asset: "XCPC", Quantity: 100, Address: "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9", Memo: "this is the test"}
	tx := xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Send_{&prototx},
	}

	fmt.Printf("%v", tx.GetMsgtype())
	// plaintx is the plain protobuf message
	plaintx, err := proto.Marshal(&tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plain protobuf sample msg:\n%x\n", plaintx)

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
	fmt.Printf("RC4 encrypted msg:\n%x\n.\n", enctx)
}

// Experiment2 test
func Experiment2() {
	var enctx []byte
	var hexkey []byte

	prototx := &xm.XCPCTransaction_Broadcast{Text: "XCPC", Value: 100, Feefraction: 1000, Timestamp: "27-May-2018"}
	tx := &xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Broadcast_{prototx},
	}
	// plaintx is the plain protobuf message
	plaintx, err := proto.Marshal(tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plain protobuf sample msg:\n%x\n", plaintx)
	fmt.Printf("Plain protobuf sample msg:\n%v\n", plaintx)

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
	fmt.Printf("RC4 encrypted msg:\n%x\n.\n", enctx)
	fmt.Printf("RC4 encrypted msg:\n%v\n.\n", enctx)
}

// ProtoOneof test the oneof module of protobuf
func ProtoOneof() {
	// This function runs the test for protobuf dealing with different transaction
	// type of messages.
	tx_broadcast := &xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Broadcast_{
			&xm.XCPCTransaction_Broadcast{
				Text:        "broadcast msg",
				Value:       100000000,
				Feefraction: 100000,
				Timestamp:   "02-June-2018",
			},
		},
	}
	tx_send := &xm.XCPCTransaction{
		Msgtype: &xm.XCPCTransaction_Send_{
			&xm.XCPCTransaction_Send{
				Asset:    "XCPC",
				Quantity: 1000,
				Address:  "Test_address",
				Memo:     "1000 XCPC TO BUY PIZZA",
			},
		},
	}
	response(tx_broadcast)
	response(tx_send)
}

func response(m *xm.XCPCTransaction) {
	switch x := m.Msgtype.(type) {
	case *xm.XCPCTransaction_Broadcast_:
		fmt.Println("broadcase message")
	case *xm.XCPCTransaction_Send_:
		fmt.Println("send message")
	case nil:
		fmt.Println("no message assigned")
	default:
		fmt.Println("Default is%T", x)
	}
}
