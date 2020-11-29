package main

import (
	"encoding/base64"
	"fmt"

	"github.com/julianlee107/learning/encryption/encyto/aesCrypto"
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
	if err != nil {
		panic(err)
	}
	anoterStr, err := aesCrypto.AESDecrypt(resultByte, Key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(anoterStr))
}
