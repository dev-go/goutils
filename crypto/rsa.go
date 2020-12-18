// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GenerateRSAKey : generate RSA key pair
func GenerateRSAKey(size int) ([]byte, []byte, error) {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, nil, err
	}
	var privateDER = x509.MarshalPKCS1PrivateKey(privateKey)
	var privateBlock = &pem.Block{
		Type:  "RSA private key",
		Bytes: privateDER,
	}
	var privateBuf = bytes.NewBuffer([]byte{})
	if err = pem.Encode(privateBuf, privateBlock); err != nil {
		return nil, nil, err
	}

	// generate public key
	var publicKey = &privateKey.PublicKey
	publicDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	var publicBlock = &pem.Block{
		Type:  "RSA public key",
		Bytes: publicDER,
	}
	var publicBuf = bytes.NewBuffer([]byte{})
	if err = pem.Encode(publicBuf, publicBlock); err != nil {
		return nil, nil, err
	}
	return privateBuf.Bytes(), publicBuf.Bytes(), nil
}

// RSAEncrypt : RSA encrypt with public key
func RSAEncrypt(publicKey []byte, originData []byte) ([]byte, error) {
	var block, _ = pem.Decode(publicKey)
	if block == nil {
		return make([]byte, 0), ErrParam
	}
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return make([]byte, 0), ErrParam
	}
	return rsa.EncryptPKCS1v15(rand.Reader, public.(*rsa.PublicKey), originData)
}

// RSADecrypt : RSA decrypt with private key
func RSADecrypt(privatekey []byte, cipher []byte) ([]byte, error) {
	var block, _ = pem.Decode(privatekey)
	if block == nil {
		return make([]byte, 0), ErrParam
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return make([]byte, 0), ErrParam
	}
	return rsa.DecryptPKCS1v15(rand.Reader, private, cipher)
}
