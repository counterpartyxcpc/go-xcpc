package message

import rc "crypto/rc4"

type XCPCMessage struct {
}

// Rc4Enc Encrypt the byte message using RC4 encryption
func Rc4Enc(key, mes []byte) ([]byte, error) {
	c, err := rc.NewCipher(key)
	ciphertext := make([]byte, len(mes))
	c.XORKeyStream(ciphertext, mes)
	c.Reset()
	return ciphertext, err
}
