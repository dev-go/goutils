// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"encoding/hex"
	"testing"
)

func TestSHA(t *testing.T) {
	var src = []byte("Hello, World !!!")
	var file = "LICENSE"

	// sha1 hash
	r1, err := SHA1Hash(src)
	t.Logf("sha1 hash: r1 = %v, err = %v", r1, err)
	if err != nil {
		t.Fatalf("sha1 hash failed: %v", err)
	}
	if hex.EncodeToString(r1) != "d7298a5154ef7da372931e3828b76ef107a7a80c" {
		t.Fatalf("sha1 hash verify failed: %v", ErrVerify)
	}
	t.Logf("sha1 hash verify: ok")

	// sha1 hash file
	r2, err := SHA1HashFile(file)
	t.Logf("sha1 hash file: r2 = %v, err = %v", r2, err)
	if err != nil {
		t.Fatalf("sha1 hash file failed: %v", err)
	}
	if hex.EncodeToString(r2) != "84fef3b6bdb88654a92fc214780c7b2cffdbeb4d" {
		t.Fatalf("sha1 hash file verify failed: %v", ErrVerify)
	}
	t.Logf("sha1 hash file verify: ok")

	// sha256 hash
	r3, err := SHA256Hash(src)
	t.Logf("sha256 hash: r3 = %v, err = %v", r3, err)
	if err != nil {
		t.Fatalf("sha256 hash failed: %v", err)
	}
	if hex.EncodeToString(r3) != "b2cfa9166dc19e24fe6a12f1c03289f827a3d2c89f928b428a2546a58c2c29d9" {
		t.Fatalf("sha256 hash verify failed: %v", ErrVerify)
	}
	t.Logf("sha256 hash verify: ok")

	// sha256 hash file
	r4, err := SHA256HashFile(file)
	t.Logf("sha256 hash file: r4 = %v, err = %v", r4, err)
	if err != nil {
		t.Fatalf("sha256 hash file failed: %v", err)
	}
	if hex.EncodeToString(r4) != "a2823636115acf09c753635c0f0b10c9eea10d7cef58488a0216bdc1867b6174" {
		t.Fatalf("sha256 hash file verify failed: %v", ErrVerify)
	}
	t.Logf("sha256 hash file verify: ok")

	// sha512 hash
	r5, err := SHA512Hash(src)
	t.Logf("sha512 hash: r5 = %v, err = %v", r5, err)
	if err != nil {
		t.Fatalf("sha512 hash failed: %v", err)
	}
	if hex.EncodeToString(r5) != "c3a5be66a4e12a9754f9d970f29c7a0d42f2d7a602706e6cd9dccbed6cbf5bedb24471c612ba8b1a0731cc37b8d816ebd18c1355510be80c354f4588276305db" {
		t.Fatalf("sha512 hash verify failed: %v", ErrVerify)
	}
	t.Logf("sha512 hash verify: ok")

	// sha512 hash file
	r6, err := SHA512HashFile(file)
	t.Logf("sha512 hash file: r6 = %v, err = %v", r6, err)
	if err != nil {
		t.Fatalf("sha512 hash file failed: %v", err)
	}
	if hex.EncodeToString(r6) != "220431e4647f0a7bc39b77beaf4c25162a60228ee4c097ecf0d9706f7b4b997189210c87eeed44573e7aa93dbe3864d5ff44501273e33b9e0893eb65bc752772" {
		t.Fatalf("sha512 hash file verify failed: %v", ErrVerify)
	}
	t.Logf("sha512 hash file verify: ok")
}
