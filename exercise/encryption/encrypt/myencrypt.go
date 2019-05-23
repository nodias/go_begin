package encrypt

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func ToHash(message string) (digest []byte) {
	hash := md5.New()
	hash.Write([]byte(message))
	digest = hash.Sum(nil)
	return
}

func SingMessage(privateKey *rsa.PrivateKey, digest []byte) (signature []byte, err error) {
	var hash crypto.Hash
	signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hash, digest)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func VertifyMessage(publicKey *rsa.PublicKey, signature []byte, digest []byte, message string) (err error) {
	var hash crypto.Hash
	err = rsa.VerifyPKCS1v15(publicKey, hash, digest, signature)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func RsaKeyGenerator() (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	publicKey = &privateKey.PublicKey
	return
}

func RsaKeyEncode(publicKey *rsa.PublicKey, plainText string) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ciphertext, nil
}

func RsaKeyDecode(privateKey rsa.PrivateKey, ciphertext []byte) (plainText string, err error) {
	plainTextByte, err := rsa.DecryptPKCS1v15(rand.Reader, &privateKey, ciphertext)
	plainText = string(plainTextByte)
	return
}
