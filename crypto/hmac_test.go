// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	stdcrypto "crypto"
	"encoding/hex"
	"testing"
)

func TestHMAC(t *testing.T) {
	var src = []byte("hello, world")
	var key = []byte("private")

	// hmac hash
	r1, err := HMACHash(stdcrypto.MD5, key, src)
	t.Logf("hmac hash: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Fatalf("hmac hash failed: %v", err)
	}
	if hex.EncodeToString(r1) != "e044378aae36ee1a705ff1061dcc3d8a" {
		t.Fatalf("hmac hash verify failed: %v", ErrVerify)
	}
	t.Logf("hmac hash verify: ok")

	// hmac verify
	if err = HMACVerify(stdcrypto.MD5, key, src, r1); err != nil {
		t.Fatalf("hmac verify failed: %v", err)
	}
	t.Logf("hmac verify: ok")
}
