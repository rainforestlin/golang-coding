package main

import (
	"encoding/base64"
	"fmt"
	"golang-coding/encryption/encyto/aesCrypto"
)

func main() {
	Key := []byte("slice bounds outdafdafda")
	Data := []byte("goroutine 1 [running]:goroutine 1 [running]:goroutine 1 [running]:goroutine 1 [running]:goroutine 1 [running]:goroutine 1 [running]:")
	result, err := aesCrypto.AESEncrypt(Data, Key)
	if err != nil {
		panic(err)
	}
	decrypto, err := aesCrypto.AESDecrypt(result, Key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decrypto))
	resultStr := aesCrypto.Base64AES(result)
	fmt.Println(resultStr)
	resultByte, err := base64.StdEncoding.DecodeString(resultStr)
	anoterStr, err := aesCrypto.AESDecrypt(resultByte, Key)
	fmt.Println(string(anoterStr))
}
