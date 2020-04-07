package aesCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//使用PKCS7填充方式
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

func PKCS7UnPadding(originData []byte) []byte {
	length := len(originData)
	upadding := int(originData[length-1])
	return originData[:(length - upadding)]
}

func AESEncrypt(originData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originData = PKCS7Padding(originData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return crypted, nil
}

func AESDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originData := make([]byte, len(crypted))
	blockMode.CryptBlocks(originData, crypted)
	originData = PKCS7UnPadding(originData)
	return originData, nil
}

func Base64AES(aesByte []byte) string {
	return base64.StdEncoding.EncodeToString(aesByte)
}
