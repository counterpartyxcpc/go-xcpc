package message

import (
	"testing"
)

type rc4TestMes struct {
	key, keystream []byte
}

type keyplainciper struct {
	key, plaintext, ciphertext []byte
}

// Test vectors from the Wikipedia page: http://en.wikipedia.org/wiki/RC4
var rc4testcases = []keyplainciper{
	{[]byte("Key"), []byte("Plaintext"), []byte{0xBB, 0xF3, 0x16, 0xE8, 0xD9, 0x40, 0xAF, 0x0A, 0xD3}},
	{[]byte("Wiki"), []byte("pedia"), []byte{0x10, 0x21, 0xBF, 0x04, 0x20}},
	{[]byte("Secret"), []byte("Attack at dawn"), []byte{0x45, 0xA0, 0x1F, 0x64, 0x5F, 0xC3, 0x5B, 0x38, 0x35, 0x52, 0x54, 0x4B, 0x9B, 0xF5}},
}

func TestXcpcRc4Enc(t *testing.T) {
	for _, kp := range rc4testcases {
		cipher := XcpcRc4Enc(kp.key, kp.plaintext)
		for i, v := range cipher {
			if v != kp.ciphertext[i] {
				t.Fatalf("test %s fail:\nmismatch at byte %d:\nhave %x\nwant %x", kp.plaintext, i, cipher, kp.ciphertext)
			}
		}
	}
}
