// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func TestRSA(t *testing.T) {
	var size = 1024
	privateKey, publicKey, err := GenerateRSAKey(size)
	if err != nil {
		t.Fatalf("generate rsa key failed: %v", err)
	}
	t.Logf("generate rsa key: ok")
	// print details
	t.Logf("generate rsa key: private_key \n%s", privateKey)
	t.Logf("generate rsa key: public_key \n%s", publicKey)

	var data = []byte("Hello, World !!!")
	cipher, err := RSAEncrypt(publicKey, data)
	if err != nil {
		t.Fatalf("rsa encrypt failed: %v", err)
	}
	t.Logf("rsa encrypt: cipher \n%s", hex.EncodeToString(cipher))
	data2, err := RSADecrypt(privateKey, cipher)
	if err != nil {
		t.Fatalf("rsa decrypt failed: %v", err)
	}
	t.Logf("rsa decrypt: data2 \n%s", data2)
	if hex.EncodeToString(data) != hex.EncodeToString(data2) {
		t.Fatalf("rsa decrypt verify failed: %v", ErrVerify)
	}
	t.Fatalf("rsa decrypt verify: ok")
}
