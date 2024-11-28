package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math"
	"math/big"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomInt64() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return n
}

func GenerateRandomInt() int {
	var n int
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return n
}

func GenerateIntUpTo(n int) int {
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return n
}
