// Copyright (c) 2020, devgo.club
// All rights reserved.

package crypto

import (
	"bufio"
	"crypto/md5"
	"io"
	"os"
)

// MD5Hash : calculate message digest with MD5
func MD5Hash(src []byte) ([]byte, error) {
	var h = md5.New()
	if _, err := h.Write(src); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}

// MD5HashFile : calculate message digest of file withc MD5
func MD5HashFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()
	var b = bufio.NewReader(f)
	var h = md5.New()
	if _, err = io.Copy(h, b); err != nil {
		return make([]byte, 0), err
	}
	return h.Sum(nil), nil
}
