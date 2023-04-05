package aesCustom

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var CIPHER_KEY = "ADDONSINTEGRATIONSIT000000000000"

type CryptoService struct {
	gcm        cipher.AEAD
	cipher_key string
}

func NewCryptoService(key []byte, cipher_key string) (cs *CryptoService, err error) {
	if len(key) < 1 {
		key = []byte(cipher_key)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &CryptoService{
		gcm:        gcm,
		cipher_key: cipher_key,
	}, nil
}

func (cs CryptoService) Encrypt(plain []byte) (cipher []byte, err error) {
	nonce := make([]byte, cs.gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	cipher = cs.gcm.Seal(nonce, nonce, plain, nil)
	cipher = append(nonce, cipher...)

	return cipher, nil
}

func (cs CryptoService) EncryptBase64(plain []byte) (base64Cipher []byte, err error) {
	cipher, err := cs.Encrypt(plain)
	if err != nil {
		return nil, err
	}

	base64Cipher = make([]byte, base64.StdEncoding.EncodedLen(len(cipher)))
	base64.StdEncoding.Encode(base64Cipher, cipher)

	return
}

func (cs CryptoService) Decrypt(cipher []byte) (plain []byte, err error) {
	nonce := cipher[0:cs.gcm.NonceSize()]
	ciphertext := cipher[cs.gcm.NonceSize():]

	plain, err = cs.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (cs CryptoService) DecryptBase64(base64Cipher []byte) (plain []byte, err error) {
	cipher := make([]byte, base64.StdEncoding.DecodedLen(len(base64Cipher)))
	_, err = base64.StdEncoding.Decode(cipher, base64Cipher)
	if err != nil {
		return nil, err
	}

	return cs.Decrypt(cipher)
}
