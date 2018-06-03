package message

import rc "crypto/rc4"
import xm "go-xcpc/proto"
import "github.com/golang/protobuf/proto"

// XCPCMessage is the wrapper structure that contain main transaction body of message
type XCPCMessage struct {
	tx     *xm.XCPCTransaction
	rc4key []byte
}

// XCPCMessageMethods contains basic methods of XCPC message
type XCPCMessageMethods interface {
	create() []byte
	//send()
}

// Rc4Enc Encrypt the byte message using RC4 encryption
func Rc4Enc(key, mes []byte) ([]byte, error) {
	c, err := rc.NewCipher(key)
	ciphertext := make([]byte, len(mes))
	c.XORKeyStream(ciphertext, mes)
	c.Reset()
	return ciphertext, err
}

func (m XCPCMessage) create() ([]byte, error) {
	// create() generate byte array from XCPCmessage
	var enctx []byte
	// prefix of xcp message
	var prefix = []byte{0x00, 0x58, 0x43, 0x50}
	plaintx, err := proto.Marshal(m.tx)
	if err != nil {
		panic(err)
	}
	encmsg, err := Rc4Enc(m.rc4key, plaintx)
	if err != nil {
		panic(err)
	}
	enctx = append(enctx, prefix...)
	enctx = append(enctx, encmsg...)
	return enctx, err
}
