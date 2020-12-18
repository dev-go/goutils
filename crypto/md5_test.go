// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func TestMD5(t *testing.T) {
	var src = []byte("hello")
	var file = "md5.go"

	// md5 hash
	r1, err := MD5Hash(src)
	t.Logf("md5 hash: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Fatalf("md5 hash failed: %v", err)
	}
	if hex.EncodeToString(r1) != "5d41402abc4b2a76b9719d911017c592" {
		t.Fatalf("md5 hash verify failed: %v", ErrVerify)
	}
	t.Logf("md5 hash verify: ok")

	// md5 hash file
	r2, err := MD5HashFile(file)
	t.Logf("md5 hash file: r2 = %v, err = %v", r2, err)
	if err != nil {
		t.Fatalf("md5 hash file failed: %v", err)
	}
	if hex.EncodeToString(r2) != "85a33b279cab137a233976091dd32aba" {
		t.Fatalf("md5 hash file verify failed: %v", ErrVerify)
	}
	t.Logf("md5 hash file verify: ok")
}
