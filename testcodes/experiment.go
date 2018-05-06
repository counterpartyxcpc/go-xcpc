package playground

import (
	"encoding/json"
	"fmt"
	enc "go-xcpc/message"
	xm "go-xcpc/proto"
	"io/ioutil"
	"os"

	proto "github.com/golang/protobuf/proto"
)

var key = []byte("1789916b3326647d973606cdbd15c4aa127e4b7079bccf599905e14a9f22593a")

// JSONToXcpcTx is a test function to load a json file and convert to protobuf structure
func JSONToXcpcTx(filename string) *xm.XCPCsend {
	var xcpcmes *xm.XCPCsend
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, &xcpcmes)
	if err != nil {
		panic(err)
	}
	return xcpcmes
}

func initxcpctx() *xm.XCPCsend {
	myxt := xm.XCPCsend{Asset: "XCPC", Quantity: 100, Address: "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9", Memo: "this is the test"}
	return &myxt
}

// XcpcToByte is a test function to load protobuf message from file
func XcpcToByte(filename string) {
	var xcpc xm.XCPCsend
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(f)
	fmt.Println()
	err = proto.Unmarshal(f, &xcpc)
	if err != nil {
		panic(err)
	}
	fmt.Println("XcpcToByte output")
	fmt.Printf("%v", xcpc)
}

// Experiment test
func Experiment() {
	examtx := initxcpctx()
	x, err := proto.Marshal(examtx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plain protobuf sample msg:\n%v\n", x)
	f, _ := os.Create("./x.log")
	defer f.Close()
	encmsg, err := enc.Rc4Enc(key, x)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RC4 encrypted msg:\n%x\n.\n", encmsg)
	f.Write(x)
	XcpcToByte("./x.log")
}
