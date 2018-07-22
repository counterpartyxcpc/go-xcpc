package encode

import (
	"crypto/rc4"

	xc "github.com/CounterpartyXCPC/go-xcpc/protobuf"
	"github.com/golang/protobuf/proto"
)

// TypeError indicate the message is not a valid xcpc_tx message
type TypeError struct {
	Msg string
}

func (e TypeError) Error() string {
	return e.Msg
}

// Message is the wrapper structure that contain main transaction body of message
type Message struct {
	txproto *xc.XCPCTransaction
	txbyte  []byte
	rc4key  []byte
}

// Rc4Enc Encrypt the byte message using RC4 encryption
func Rc4Enc(key, mes []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	ciphertext := make([]byte, len(mes))
	c.XORKeyStream(ciphertext, mes)
	c.Reset()
	return ciphertext, err
}

// Serialize generate byte array from XCPCmessage structure
func (m *Message) Serialize() error {
	// serialize() generate byte array from XCPCmessage
	var enctx []byte
	// prefix of xcp message
	var prefix = []byte{0x00, 0x58, 0x43, 0x50}
	plaintx, err := proto.Marshal(m.txproto)
	if err != nil {
		return err
	}
	encmsg, err := Rc4Enc(m.rc4key, plaintx)
	if err != nil {
		return err
	}
	enctx = append(enctx, prefix...)
	enctx = append(enctx, encmsg...)
	m.txbyte = enctx
	return err
}

//Load create XCPCMessage structure from byte array
func (m *Message) Load(b []byte) error {
	var prefix = []byte{0x00, 0x58, 0x43, 0x50}
	for i, v := range prefix {
		if b[i] == v {
			continue
		} else {
			return TypeError{"Not a valid XCPC message, mismatch XCPC prefix"}
		}
	}
	m.txbyte = b
	return nil
}
