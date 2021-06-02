package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type Crypter struct {
	gcm cipher.AEAD
}

func (cp Crypter) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, cp.gcm.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return cp.gcm.Seal(nonce, nonce, data, nil), err
}

func (cp Crypter) Decrypt(data []byte) ([]byte, error) {
	nonceSize := cp.gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("wrong data")
	}
	nonce, data := data[:nonceSize], data[nonceSize:]
	decrypted, err := cp.gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func New(key []byte) (*Crypter, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	cp := &Crypter{
		gcm: gcm,
	}
	return cp, nil
}
