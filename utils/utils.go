// Copyright (C) 2021-2023 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package utils

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"os"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ReadAllText(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}

func WriteAllText(path, text string) {
	_ = os.WriteFile(path, []byte(text), 0644)
}

func SliceXor(s1, s2 []byte) (res []byte) {
	for k, b := range s1 {
		newBt := b ^ s2[k]
		res = append(res, newBt)
	}

	return
}

func Sub16(str []byte) []byte {
	return str[16:]
}

func RandomStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	mRand := make([]byte, len)
	rand.Read(mRand)
	mRandStr := hex.EncodeToString(mRand)[0:len]

	return mRandStr
}

func RandomBase64Token(len int) string {
	rand.NewSource(time.Now().UnixNano())
	mRand := make([]byte, len)
	rand.Read(mRand)
	mRandStr := hex.EncodeToString(mRand)[0:len]

	return base64.RawStdEncoding.EncodeToString([]byte(mRandStr))
}