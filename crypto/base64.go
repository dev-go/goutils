// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bytes"
	"encoding/base64"
)

var cutset = string([]byte{'\x00'})

// Base64EncodeSTD : BASE64 encode
func Base64EncodeSTD(src []byte) []byte {
	var dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// Base64DecodeSTD : BASE64 decode
func Base64DecodeSTD(src []byte) ([]byte, error) {
	var dst = make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	if _, err := base64.StdEncoding.Decode(dst, src); err != nil {
		return make([]byte, 0), err
	}
	return bytes.TrimRight(dst, cutset), nil
}

// Base64EncodeURL : BASE64 encode for url
func Base64EncodeURL(src []byte) []byte {
	var dst = make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(dst, src)
	return dst
}

// Base64DecodeURL : BASE64 decode for url
func Base64DecodeURL(src []byte) ([]byte, error) {
	var dst = make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	if _, err := base64.URLEncoding.Decode(dst, src); err != nil {
		return make([]byte, 0), err
	}
	return bytes.TrimRight(dst, cutset), nil
}
