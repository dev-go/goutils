// Copyright (c) 2020, devgo.club
// All rights reserved.

package aescrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"math/rand"
)

const encodeME = "ABCTUyz0123YZaLMNOVWXstuvwxPQibcdDEFGHIJKjklmnopqRSefghr456789-_"

var meEncoding = base64.NewEncoding(encodeME)

// errors
var (
	ErrEncrypto = errors.New("encrypto failed")
	ErrDecrypto = errors.New("decrypto failed")
)

// Encrypt : 加密
func Encrypt(msg string, key string) (string, error) {
	var block, err = aes.NewCipher(paddingKey(key))
	if err != nil {
		return "", ErrEncrypto
	}
	var blockSize = block.BlockSize()
	var r = byte(rand.Int() & 0xff)
	var _msg = append([]byte(msg), r)
	var origData = pkcs5Padding(_msg, blockSize)
	var crc = crc32.ChecksumIEEE(origData)
	var iv = genIV(blockSize)
	var blockMode = cipher.NewCBCEncrypter(block, iv)
	var crypted = make([]byte, len(origData)+6)
	crypted[4], crypted[5] = iv[0], iv[1]
	blockMode.CryptBlocks(crypted[6:], origData)
	binary.BigEndian.PutUint32(crypted[:4], crc)
	return meEncoding.EncodeToString(crypted), nil
}

func paddingKey(key string) []byte {
	var newkey = key + "01234567890123456789"
	return []byte(newkey)[:16]
}

func genIV(size int) []byte {
	var r = rand.Int()
	var buf = make([]byte, size)
	for i := range buf {
		if i%2 == 0 {
			buf[i] = byte(r & 0xff)
		} else {
			buf[i] = byte((r >> 8) & 0xff)
		}
	}
	return buf
}

// Decrypt : 解密
func Decrypt(msg string, key string) (string, error) {
	crypted, err := meEncoding.DecodeString(msg)
	if err != nil {
		return "", ErrDecrypto
	}
	if len(crypted) <= 7 {
		return "", ErrDecrypto
	}
	var crc = binary.BigEndian.Uint32(crypted[:4])
	var block cipher.Block
	block, err = aes.NewCipher(paddingKey(key))
	if err != nil {
		return "", ErrDecrypto
	}
	var blockSize = block.BlockSize()
	var iv = make([]byte, blockSize)
	for i := range iv {
		if i%2 == 0 {
			iv[i] = crypted[4]
		} else {
			iv[i] = crypted[5]
		}
	}
	var origData = make([]byte, len(crypted)-6)
	var blockMode = cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(origData, crypted[6:])
	if crc != crc32.ChecksumIEEE(origData) {
		return "", ErrDecrypto
	}
	origData = pkcs5Unpadding(origData)
	return string(origData[0 : len(origData)-1]), nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
