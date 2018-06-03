package message

import (
	rc "crypto/rc4"
	xm "go-xcpc/proto"

	"github.com/golang/protobuf/proto"
)

// XCPCMessage is the wrapper structure that contain main transaction body of message
type XCPCMessage struct {
	txproto *xm.XCPCTransaction
	txbyte  []byte
	rc4key  []byte
}

// Rc4Enc Encrypt the byte message using RC4 encryption
func Rc4Enc(key, mes []byte) ([]byte, error) {
	c, err := rc.NewCipher(key)
	ciphertext := make([]byte, len(mes))
	c.XORKeyStream(ciphertext, mes)
	c.Reset()
	return ciphertext, err
}

func (m XCPCMessage) create() error {
	// create() generate byte array from XCPCmessage
	var enctx []byte
	// prefix of xcp message
	var prefix = []byte{0x00, 0x58, 0x43, 0x50}
	plaintx, err := proto.Marshal(m.txproto)
	if err != nil {
		panic(err)
	}
	encmsg, err := Rc4Enc(m.rc4key, plaintx)
	if err != nil {
		panic(err)
	}
	enctx = append(enctx, prefix...)
	enctx = append(enctx, encmsg...)
	m.txbyte = enctx
	return err
}

func (m XCPCMessage) loadxcpc(b []byte) error {
	var prefix = []byte{0x00, 0x58, 0x43, 0x50}
	for i, e := range prefix {
		if b[i] == e {
			continue
		} else {
			continue
		}
	}
}
