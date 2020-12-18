// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func Test3DES(t *testing.T) {
	var src = []byte("hello, world! I am a gopher.")
	var key = []byte("go-utils, from a gopher.") // CAUTION: the length of key must be 24!!!

	// 3des encrypt
	r1, err := TripleDESEncrypt(src, key, PKCS5PaddingMode)
	t.Logf("3des encrypt: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Fatalf("3des encrypt failed: %v", err)
	}
	if hex.EncodeToString(r1) != `267a85269451aa3f830f4dacb02e9f9f0a113a22c3900e08b9e05e98866d96ed` {
		t.Fatalf("3des encrypt verify failed: %v", ErrVerify)
	}
	t.Logf("3des encrypt verify: ok")

	// 3des decrypt
	r2, err := TripleDESDecrypt(r1, key, PKCS5PaddingMode)
	t.Logf("3des decrypt: r2 = %v, err = %v", r2, err)
	if err != nil {
		t.Fatalf("3des decrypt failed: %v", err)
	}
	if hex.EncodeToString(r2) != hex.EncodeToString(src) {
		t.Fatalf("3des decrypt verify failed: %v", ErrVerify)
	}
	t.Logf("3des decrypt verify: ok")
}
