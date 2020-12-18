// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func TestDES(t *testing.T) {
	var src = []byte("hello, world! 你好，世界！")
	var key = []byte("go-utils") // CAUTION: the length of key must be 8!!!

	// des encrypt with pkcs5 padding
	r1, err := DESEncrypt(src, key, PKCS5PaddingMode)
	t.Logf("des encrypt: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Fatalf("des encrypt failed: %v", err)
	}
	if hex.EncodeToString(r1) != `40cc6acce47497eceb9ad6479383b7f983b080e5b047896aa046123f2a9a12f5b02d1e23c89afc19` {
		t.Fatalf("des encrypt verify failed: %v", ErrVerify)
	}
	t.Logf("des encrypt verify: ok")

	// des decrypt with pkcs5 unpadding
	r2, err := DESDecrypt(r1, key, PKCS5PaddingMode)
	t.Logf("des decrypt: r2 = %v, err = %v", r2, err)
	if err != nil {
		t.Fatalf("des decrypt failed: %v", err)
	}
	if hex.EncodeToString(r2) != hex.EncodeToString(src) {
		t.Fatalf("des decrypt verify failed: %v", ErrVerify)
	}
	t.Logf("des decrypt verify: ok")

	// des encrypt with zeros padding
	r3, err := DESEncrypt(src, key, ZerosPaddingMode)
	t.Logf("des encrypt: r3 = %v, err = %v", r3, err)
	if err != nil {
		t.Fatalf("des encrypt failed: %v", err)
	}
	if hex.EncodeToString(r3) != `40cc6acce47497eceb9ad6479383b7f983b080e5b047896aa046123f2a9a12f5b7e0695b189d51d3` {
		t.Fatalf("des encrypt verify failed: %v", ErrVerify)
	}
	t.Logf("des encrypt verify: ok")

	// des decrypt with zeros unpadding
	r4, err := DESDecrypt(r3, key, ZerosPaddingMode)
	t.Logf("des decrypt: r4 = %v, err = %v", r4, err)
	if err != nil {
		t.Fatalf("des decrypt failed: %v", err)
	}
	if hex.EncodeToString(r4) != hex.EncodeToString(src) {
		t.Fatalf("des decrypt verify failed: %v", ErrVerify)
	}
	t.Logf("des decrypt verify: ok")
}
