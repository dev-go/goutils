// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESEncrypt : AES encrypt with CBC mode
func AESEncrypt(src []byte, key []byte, padding paddingMode) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}
	var blockMode = cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	paddingFunc, err := GetPadding(padding)
	if err != nil {
		return make([]byte, 0), err
	}
	src = paddingFunc(src, block.BlockSize())
	var result = make([]byte, len(src))
	blockMode.CryptBlocks(result, src)
	return result, nil
}

// AESDecrypt : AES decrypt with CBC mode
func AESDecrypt(src []byte, key []byte, padding paddingMode) ([]byte, error) {
	block, err := aes.NewCipher(key)
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
