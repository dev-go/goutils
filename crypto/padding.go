// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bytes"
)

// padding modes
const (
	ZerosPaddingMode = paddingMode(`Zeros`)
	PKCS5PaddingMode = paddingMode(`PKCS5`)
)

type paddingMode string

// GetPadding : get padding function
func GetPadding(mode paddingMode) (func([]byte, int) []byte, error) {
	switch mode {
	case ZerosPaddingMode:
		return zerosPadding, nil
	case PKCS5PaddingMode:
		return pkcs5Padding, nil
	default:
		return nil, ErrParam
	}
}

// GetUnpadding : get unpadding function
func GetUnpadding(mode paddingMode) (func([]byte) []byte, error) {
	switch mode {
	case ZerosPaddingMode:
		return zerosUnpadding, nil
	case PKCS5PaddingMode:
		return pkcs5Unpadding, nil
	default:
		return nil, ErrParam
	}
}

// zerosPadding : zeros padding mode
func zerosPadding(cipherText []byte, blockSize int) []byte {
	var paddingCount = blockSize - len(cipherText)%blockSize
	var paddingText = bytes.Repeat([]byte{0}, paddingCount)
	return append(cipherText, paddingText...)
}

// zerosUnpadding : zeros unpadding mode
func zerosUnpadding(originData []byte) []byte {
	return bytes.TrimRightFunc(originData, func(r rune) bool {
		return r == rune(0)
	})
}

// pkcs5Padding : PKCS5 padding mode
func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	var paddingCount = blockSize - len(cipherText)%blockSize
	var paddingText = bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(cipherText, paddingText...)
}

// pkcs5Unpadding : PKCS5 unpadding mode
func pkcs5Unpadding(originData []byte) []byte {
	var length = len(originData)
	var count = int(originData[length-1])
	return originData[:(length - count)]
}
