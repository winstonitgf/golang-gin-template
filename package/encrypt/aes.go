package encrypt

import (
	"bao-bet365-api/package/log"
	"crypto/aes"
	"crypto/cipher"
)

var (
	key = []byte("P~sp)Sx(#0W,`v!K~vb)Tb(#0A,`v!QQ")
)

func AesEncrypt(data string) string {
	origData := []byte(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	ff := encodeBase64(crypted)
	return ff
}

func AesDecrypt(data string) (string, error) {
	crypted := decodeBase64(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}
