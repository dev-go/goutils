// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func TestAES(t *testing.T) {
	var src = []byte(`{"user":"Terrence", "expire":12000}`)
	var key = []byte("hello,gohello,go") // length: 16 --> AES-128

	// aes encrypt
	r1, err := AESEncrypt(src, key, PKCS5PaddingMode)
	t.Logf("aes encrypt: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Logf("aes encrypt failed: %v", err)
	}
	if hex.EncodeToString(r1) != `96f649b434b7a6d1e1c2a8df492b60d6f34d86056b29b16cd3331d81686a82a6d59a071693811b0ea77f39209de6e56f` {
		t.Logf("aes encrypt verify failed: %v", ErrVerify)
	}
	t.Logf("aes encrypt verify: ok")

	// aes decrypt
	r2, err := AESDecrypt(r1, key, PKCS5PaddingMode)
	t.Logf("aes decrypt: r2 = %v, err = %v", r2, err)
	if err != nil {
		t.Logf("aes decrypt failed: %v", err)
	}
	if hex.EncodeToString(r2) != hex.EncodeToString(src) {
		t.Logf("aes decrypt verify failed: %v", ErrVerify)
	}
	t.Logf("aes decrypt verify: ok")
}
