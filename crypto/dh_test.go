// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"testing"
)

func TestDH(t *testing.T) {
	comm, err := GetDHPublic(14)
	if err != nil {
		t.Fatalf("get dh public failed: %v", err)
	}
	one, err := GenerateKey(comm, nil)
	if err != nil {
		t.Fatalf("generate key failed: %v", err)
	}
	two, err := GenerateKey(comm, nil)
	if err != nil {
		t.Fatalf("generate key failed: %v", err)
	}
	oneKey, err := CalculateExchangeKey(comm, one, two)
	if err != nil {
		t.Fatalf("calculate exchange key failed: %v", err)
	}
	twoKey, err := CalculateExchangeKey(comm, two, one)
	if err != nil {
		t.Fatalf("calculate exchange key failed: %v", err)
	}
	if oneKey.String() != twoKey.String() {
		t.Fatalf("exchange key verify failed: %v", ErrVerify)
	}
	t.Logf("exchange key verify: ok")
	// // print details
	// t.Logf("++++++++++++++++++++++++++++++++++++++++++++++++++")
	// t.Logf("comm: p = %v", comm.p)
	// t.Logf("comm: g = %v", comm.g)
	// t.Logf("++++++++++++++++++++++++++++++++++++++++++++++++++")
	// t.Logf("one: public key = %v", one.y)
	// t.Logf("one: private key = %v", one.x)
	// t.Logf("++++++++++++++++++++++++++++++++++++++++++++++++++")
	// t.Logf("two: public key = %v", two.y)
	// t.Logf("two: private key = %v", two.x)
	// t.Logf("++++++++++++++++++++++++++++++++++++++++++++++++++")
	// t.Logf("one_key: %v", oneKey)
	// t.Logf("two_key: %v", twoKey)
	// t.Logf("++++++++++++++++++++++++++++++++++++++++++++++++++")
}
