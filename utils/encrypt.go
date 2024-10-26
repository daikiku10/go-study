package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AES256で暗号化する
func Aes256Encrypt(plainText, encKey, iv string) ([]byte, error) {
	bKey := make([]byte, 32) // AES-256ではキーの長さが32バイト(256ビット)である必要があるため
	copy(bKey, encKey)       // encKeyからbKeyへのデータコピーを行う [101 114 102 100 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	bIV := []byte(iv)        // authKeyをバイトスライスへ変換 [51 50 56 48 109 49 56 120 117 122 84 98 71 119 67 55]

	// 暗号キーを使用して、新しいAES暗号ブロックを生成
	block, _ := aes.NewCipher(bKey)

	plainTextBytes := []byte(plainText)
	plainTextBytes = pkcs7Pad(plainTextBytes, aes.BlockSize)

	encrypted := make([]byte, len(plainTextBytes))

	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(encrypted, plainTextBytes)

	return encrypted, nil
}

// pkcs7Pad PKCS#7パディング
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)

	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(data, padText...)
}
