// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
	"os"
)

// SHA1Hash : calculate message digest with SHA-1
func SHA1Hash(src []byte) ([]byte, error) {
	var h = sha1.New()
	var b = bytes.NewBuffer(src)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// SHA1HashFile : calculate message digest of file with SHA-1
func SHA1HashFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	var h = sha1.New()
	var b = bufio.NewReader(f)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// SHA256Hash : calculate message digest with SHA-256
func SHA256Hash(src []byte) ([]byte, error) {
	var h = sha256.New()
	var b = bytes.NewBuffer(src)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil

}

// SHA256HashFile : calculate message digest of file with SHA-256
func SHA256HashFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	var h = sha256.New()
	var b = bufio.NewReader(f)
	if _, err = io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// SHA512Hash : calculate message digest with SHA-512
func SHA512Hash(src []byte) ([]byte, error) {
	var h = sha512.New()
	var b = bytes.NewBuffer(src)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// SHA512HashFile : calculate message digest of file with SHA-512
func SHA512HashFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	var h = sha512.New()
	var b = bufio.NewReader(f)
	if _, err := io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}
