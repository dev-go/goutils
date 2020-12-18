// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"testing"
)

func TestBase64(t *testing.T) {
	var data = "你好+世界"

	// base64 std encode
	var r1 = Base64EncodeSTD([]byte(data))
	t.Logf("base64 encode: r1 = %q", r1)
	if string(r1) != "5L2g5aW9K+S4lueVjA==" {
		t.Fatalf("base64 encode verify failed: %v", ErrVerify)
	}
	t.Logf("base64 encode verify: ok")

	// base64 std decode
	r2, err := Base64DecodeSTD(r1)
	t.Logf("base64 decode: r2 = %q, err = %v", r2, err)
	if err != nil {
		t.Fatalf("base64 decode failed: %v", err)
	}
	if string(r2) != data {
		t.Fatalf("base64 decode verify failed: %v", ErrVerify)
	}
	t.Logf("base64 decode verify: ok")

	// base64 url encode
	var r3 = Base64EncodeURL([]byte(data))
	t.Logf("base64 encode for url: r3 = %q", r3)
	if string(r3) != "5L2g5aW9K-S4lueVjA==" {
		t.Fatalf("base64 encode for url verify failed: %v", ErrVerify)
	}
	t.Logf("base64 encode for url verify: ok")

	// base64 url decode
	r4, err := Base64DecodeURL(r3)
	t.Logf("base64 decode for url: r4 = %v, err = %v", r4, err)
	if err != nil {
		t.Fatalf("base64 decode for url failed: %v", err)
	}
	if string(r4) != data {
		t.Fatalf("base64 decode for url verify failed: %v", ErrVerify)
	}
	t.Logf("base64 decode for url verify: ok")
}
