package encrypt

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func DesECBEncrypt(encryptStr, encryptKey string) (string, error) {

	data := []byte(encryptStr)
	key := []byte(encryptKey)

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return encodeBase64(out), nil
}

func DesCBCEncrypt(encryptStr, encryptKey string) (string, error) {

	origData := []byte(encryptStr)
	key := []byte(encryptKey)

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return encodeBase64(crypted), nil
}
