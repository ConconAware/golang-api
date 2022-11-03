package encode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

var key_share string = os.Getenv("AES_KEY")

func SymmetricMain() {
	enc, _ := EncryptCBCMessage(secretMessage)
	fmt.Printf("Encrypt CBC %s", enc)

	dec, _ := DecryptCBCMessage(enc)
	fmt.Printf("\n Decrypt CBC %s", dec)
}

func EncryptCBCMessage(message string) (string, error) {

	if len(key_share) == 0 {
		key_share = "atshappycyptodev"
	}

	encodeBytes := []byte(message)
	//Generate ciphertext according to key
	block, err := aes.NewCipher([]byte(key_share))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	ciphertext := make([]byte, aes.BlockSize+len(encodeBytes))
	iv := ciphertext[:aes.BlockSize]

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func DecryptCBCMessage(message string) (string, error) {
	if len(key_share) == 0 {
		log.Println("use default Key")
		key_share = "atshappycyptodev"
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key_share))
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(decodeBytes))
	iv := ciphertext[:aes.BlockSize]
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
