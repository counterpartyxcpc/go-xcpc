package message

import rc "crypto/rc4"

// XcpcRc4Enc Encrypt the byte message using RC4 encryption
func XcpcRc4Enc(key, mes []byte) []byte {
	c, err := rc.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(mes))
	c.XORKeyStream(ciphertext, mes)
	c.Reset()
	return ciphertext
}
