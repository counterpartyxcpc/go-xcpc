package message

import (
	"encoding/hex"
	"fmt"
	"testing"

	xm "go-xcpc/proto"

	proto "github.com/golang/protobuf/proto"
)

type rc4TestMes struct {
	key, keystream []byte
}

type keyplainciper struct {
	key, plaintext, ciphertext []byte
}

type keyciper struct {
	key, ciphertext []byte
}

// Test vectors from the Wikipedia page: http://en.wikipedia.org/wiki/RC4
var rc4testcases = []keyplainciper{
	{[]byte("Key"), []byte("Plaintext"), []byte{0xBB, 0xF3, 0x16, 0xE8, 0xD9, 0x40, 0xAF, 0x0A, 0xD3}},
	{[]byte("Wiki"), []byte("pedia"), []byte{0x10, 0x21, 0xBF, 0x04, 0x20}},
	{[]byte("Secret"), []byte("Attack at dawn"), []byte{0x45, 0xA0, 0x1F, 0x64, 0x5F, 0xC3, 0x5B, 0x38, 0x35, 0x52, 0x54, 0x4B, 0x9B, 0xF5}},
	// counterjs testexamples, page: https://github.com/visvirial/CounterJS/blob/master/test/Message.js
	{
		[]byte{0xb3, 0x4d, 0xdf, 0x89, 0x04, 0xbc, 0xfc, 0x45, 0x4c, 0x6f, 0x06, 0xd3, 0x3e, 0x60, 0x09, 0x42, 0xc7, 0xce, 0x8f, 0x75, 0xdd, 0x2d, 0x46, 0x53, 0x2f, 0x26, 0x3a, 0x6e, 0x56, 0xd8, 0x3d, 0x34},
		[]byte{0x5e, 0x1e, 0xf3, 0xf9, 0x9e, 0x3a, 0x89, 0x06, 0x0c, 0x43, 0xca, 0xac, 0xc0, 0xa0, 0x5c, 0x15, 0x67, 0x8e, 0xf8, 0xe9, 0xba, 0x96, 0xf4, 0x2e, 0x8d, 0xc6, 0x4f, 0xa0, 0x4d, 0xda, 0x75, 0x9c, 0x2b, 0x0f, 0x4f, 0x8c, 0x34, 0xb9, 0x1a, 0xcf, 0x6f, 0x86, 0xe7},
		[]byte{0x43, 0x4e, 0x54, 0x52, 0x50, 0x52, 0x54, 0x59, 0x00, 0x00, 0x00, 0x14, 0x83, 0x22, 0x22, 0x8e, 0x65, 0x67, 0x58, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x66, 0x75, 0x67, 0x61},
	},
	{
		[]byte{0xdf, 0x6f, 0x0d, 0xe6, 0x39, 0x58, 0x07, 0x9f, 0x28, 0x4c, 0xc5, 0x12, 0xec, 0xfd, 0x84, 0x34, 0x37, 0xe4, 0x0b, 0x58, 0xbf, 0x47, 0x8b, 0x87, 0xd8, 0x64, 0xf3, 0xf5, 0xf4, 0xde, 0xe2, 0x0e},
		[]byte{0x41, 0x0a, 0x09, 0x5f, 0x92, 0x65, 0x43, 0x17, 0xca, 0x66, 0x94, 0x84, 0x13, 0x00, 0x84, 0x90, 0x92, 0xa6, 0xf1, 0x04, 0x3d, 0xdb, 0xd8, 0x98, 0x0e, 0x26, 0x9b, 0x44, 0x17, 0xcf, 0x2d, 0x2b, 0x5b, 0x0d, 0xf0, 0x3a, 0xfc, 0x16, 0x4a, 0xe8, 0x65, 0x89, 0x6e},
		[]byte{0x43, 0x4e, 0x54, 0x52, 0x50, 0x52, 0x54, 0x59, 0x14, 0x8f, 0x78, 0x71, 0x52, 0x11, 0xfd, 0xc7, 0xf8, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x76, 0x64, 0x67, 0x72, 0x75, 0x63, 0x6b},
	},
}

// This is a testcase for xcp transaction in bitcoin network
var xcprc4cases = []keyciper{
	// Transaction link: https://www.blocktrail.com/BTC/tx/949c1596236102065db47b4a164c04a8766c255a631fbf7ea93798efd4db6155
	// counterparty asset name: TRIGGERS
	// Input address first transaction ID
	// OP_RETURN
	{
		[]byte("1567798bbc628fc3ac03e74d888d5504324dc5c4dc59db16d4a336ec5ad8b701"),
		[]byte("a2f10506c0242cfcc5b0dc9204b89a6c2c2814c0610be4fbddfc3d2c14209bb679f9c538a582810b5e1f5a6b6533"),
	},
	// Transaction link: https://www.blocktrail.com/BTC/tx/933884dff9d930fdaf5ae3bf335293faa3ffa08b622b912ee7c6558029bde2a4
	// counterparty asset name: XCP
	// Input address first transaction ID
	// OP_RETURN
	{
		[]byte("2dbebbbd930613fd7ad3851ab0ff844fbbdab8c21b3df95f902fd30f8932420a"),
		[]byte("1dbff188792e64d8fe2bd6cf0429bfe2e843a07804a3f2cf308d8afc5a42650c5d1319803899768167ab94757785"),
	},
	// Transaction link: https://www.blocktrail.com/tBTC/tx/c9de23c23a2f31005b447d6f527c66f5e9917f644fa8d1172ef31d737fcf1821#tx_messages
	// counterparty asset name: XCP
	// Input address first transaction ID
	// OP_RETURN
	{
		[]byte("3e9acc6ffa8be2bd629f74f7fd369cc29f4dc207ba431888a007e10b608a7962"),
		[]byte("53c8b4c54e312d74b81ec8d07b6f3ad27ffdeb489b85ba15c2249fa40237d85b0e2ee26214d2ee04cac0c3eda771c1ceb8e7336aafb8"),
	},
	// Transaction link: https://www.blocktrail.com/tBTC/tx/eee67f9955f4cb77d38f49c4294704909adea247aaacffd93873d5032e84e214#tx_messages
	// counterparty asset name: XCP
	// send message: 10 xcp
	// Input address first transaction ID
	// OP_RETURN
	{
		[]byte("c9de23c23a2f31005b447d6f527c66f5e9917f644fa8d1172ef31d737fcf1821"),
		[]byte("483053a3223e3df277bb72831a6604de7744054b0927f7d0414a47746438ac4832bc20f3d2f59668fa721549f63e"),
	},
	// Transaction link: https://www.blocktrail.com/tBTC/tx/3e9acc6ffa8be2bd629f74f7fd369cc29f4dc207ba431888a007e10b608a7962
	{
		[]byte("e2cd21105af9f6ba542e84497633d4af6e0985d9a389f434e030542126a78a5f"),
		[]byte("b82a5a83dcab84cd13c1a69a2fc4e6585b0e4122e4e9a721305f81f4b50f705d5e3fba9a9e4cfdd384796fb36d25"),
	},
	// Transaction link: https://www.blocktrail.com/tBTC/tx/e2cd21105af9f6ba542e84497633d4af6e0985d9a389f434e030542126a78a5f
	{
		[]byte("1789916b3326647d973606cdbd15c4aa127e4b7079bccf599905e14a9f22593a"),
		[]byte("1ea10fe4057187e74949431d01f73f4892fe13a72afbc856ce42db8f6ec29b2a0a7de7fa9a68ce942bbaa5b956fdec88c6a62fbe30dafdf2"),
	},
}

var prototestcases = []xm.XCPCTransaction_Send{
	{Asset: "XCPC", Quantity: 100, Address: "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9"},
}

// TestRc4Enc verify the RC4 encryption is working properly with respect to the most common test examples
func TestRc4Enc(t *testing.T) {
	for _, kp := range rc4testcases {
		cipher, err := Rc4Enc(kp.key, kp.plaintext)
		if err != nil {
			t.Fatalf("Rc4Enc function error:\n%v", err)
		}
		for i, v := range cipher {
			if v != kp.ciphertext[i] {
				t.Fatalf("test %s fail:\nmismatch at byte %d:\nhave %x\nwant %x", kp.plaintext, i, cipher, kp.ciphertext)
			}
		}
	}
}

// TestProtoXCPCTxEncAndDec test the RC4 encryption and decryption of protobuf serialized message
func TestProtoXCPCTxEncAndDec(t *testing.T) {
	for j, msg := range prototestcases {
		b, err := proto.Marshal(&msg)
		if err != nil {
			t.Fatalf("test %d - (%v) fail:\nMarshal protobuf error:\n%v", j, msg, err)
		}
		key := []byte("Key")
		cipher, err := Rc4Enc(key, b)
		if err != nil {
			t.Fatalf("test %d - (%v) fail:\nRc4Enc function error:\n%v", j, msg, err)
		}
		// fmt.Printf("%x\n", cipher)
		plaintext, err := Rc4Enc(key, cipher)
		if err != nil {
			t.Fatalf("test %d - (%v) fail:\nRc4Enc function error:\n%v", j, msg, err)
		}
		for i, v := range plaintext {
			if v != b[i] {
				t.Fatalf("test %d - (%v) fail:\nmismatch at byte %d:\nhave %x\nwant %x", j, msg, i, plaintext, b)
			}
		}
	}
}

// Reveal the XCP message plaintext
func TestXCPRC4(t *testing.T) {
	for j, xcp := range xcprc4cases {
		kh := make([]byte, hex.DecodedLen(len(xcp.key)))
		ciph := make([]byte, hex.DecodedLen(len(xcp.ciphertext)))
		n, err := hex.Decode(kh, xcp.key)
		if err != nil {
			t.Fatalf("test XCP example %d fail:\n%x hex decode error\ndecoded bytes %d", j, xcp.key, n)
		}
		m, err := hex.Decode(ciph, xcp.ciphertext)
		if err != nil {
			t.Fatalf("test XCP example %d fail:\n%x hex decode error\ndecoded bytes %d", j, xcp.ciphertext, m)
		}
		p, err := Rc4Enc(kh, ciph)
		if err != nil {
			t.Fatalf("test XCP example %d fail:\nRc4Enc err %v", j, err)
		}
		fmt.Printf("%d\n%x\n%v\n%s\n", j, p, p, p)
	}
}

// Create byte array from XCPC message
func TestXCPCMessage(t *testing.T) {
	var xcpmsg []byte
	sendmsg := xm.XCPCTransaction_Send{
		Asset:    "XCPC",
		Quantity: 1000,
		Address:  "1EtwuGeP6t6bAJjKCHuRC67MFi4pqXF5i9",
		Memo:     "Test message",
	}
	x := XCPCMessage{
		tx: &xm.XCPCTransaction{
			Msgtype: &xm.XCPCTransaction_Send_{&sendmsg},
		},
		rc4key: []byte("1789916b3326647d973606cdbd15c4aa127e4b7079bccf599905e14a9f22593a"),
	}
	xcpmsg, err := x.create()
	if err != nil {
		t.Fatalf("TestXCPCMessage error: %v", err)
	}
	fmt.Printf("%x\n%s\n", xcpmsg, xcpmsg)
}
