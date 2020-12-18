// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"crypto/cipher"
	"crypto/des"
)

// TripleDESEncrypt : 3DES encrypt with CBC mode
func TripleDESEncrypt(src []byte, key []byte, padding paddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}
	paddingFunc, err := GetPadding(padding)
	if err != nil {
		return make([]byte, 0), err
	}
	src = paddingFunc(src, block.BlockSize())
	var blockMode = cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	var result = make([]byte, len(src))
	blockMode.CryptBlocks(result, src)
	return result, nil
}

// TripleDESDecrypt : 3DES decrypt with CBC mode
func TripleDESDecrypt(src []byte, key []byte, padding paddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}
	var blockMode = cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	var result = make([]byte, len(src))
	blockMode.CryptBlocks(result, src)
	unpaddingFunc, err := GetUnpadding(padding)
	if err != nil {
		return make([]byte, 0), err
	}
	return unpaddingFunc(result), nil
}
