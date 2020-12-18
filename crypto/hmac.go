// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bufio"
	"bytes"
	stdcrypto "crypto"
	stdhmac "crypto/hmac"
	"io"
	"os"
)

// HMAC : keyed-hash message authentication code
type HMAC struct {
	hash stdcrypto.Hash
	key  []byte
}

// NewHMAC : create a new HMAC
func NewHMAC(hash stdcrypto.Hash, key []byte) (*HMAC, error) {
	var hmac = &HMAC{
		hash: hash,
		key:  key,
	}
	if !hmac.hash.Available() {
		return nil, ErrParam
	}
	return hmac, nil
}

// HMACHash : generate message authentication code
func HMACHash(hash stdcrypto.Hash, key []byte, src []byte) ([]byte, error) {
	hmac, err := NewHMAC(hash, key)
	if err != nil {
		return make([]byte, 0), err
	}
	var h = stdhmac.New(hmac.hash.New, hmac.key)
	var b = bytes.NewBuffer(src)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// HMACVerify : verify message authentication code
func HMACVerify(hash stdcrypto.Hash, key []byte, src []byte, checksum []byte) error {
	_checksum, err := HMACHash(hash, key, src)
	if err != nil {
		return err
	}
	if !stdhmac.Equal(_checksum, checksum) {
		return ErrVerify
	}
	return nil
}

// HMACHashFile : generate message authentication code of file
func HMACHashFile(hash stdcrypto.Hash, key []byte, file string) ([]byte, error) {
	hmac, err := NewHMAC(hash, key)
	if err != nil {
		return make([]byte, 0), err
	}
	f, err := os.Open(file)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	var h = stdhmac.New(hmac.hash.New, hmac.key)
	var b = bufio.NewReader(f)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// HMACVerifyFile : generate message authentication code of file
func HMACVerifyFile(hash stdcrypto.Hash, key []byte, file string, checksum []byte) error {
	_checksum, err := HMACHashFile(hash, key, file)
	if err != nil {
		return err
	}
	if !stdhmac.Equal(_checksum, checksum) {
		return ErrVerify
	}
	return nil
}
