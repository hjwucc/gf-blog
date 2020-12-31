package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	crated := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crated, origData)

	return base64.StdEncoding.EncodeToString(crated)

}

func AesDecrypt(crated string, key string) string {
	// 转成字节数组
	cratedByte, _ := base64.StdEncoding.DecodeString(crated)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cratedByte))
	// 解密
	blockMode.CryptBlocks(orig, cratedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphered []byte, blksize int) []byte {
	padding := blksize - len(ciphered)%blksize
	padget := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphered, padget...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unfading := int(origData[length-1])
	return origData[:(length - unfading)]
}
