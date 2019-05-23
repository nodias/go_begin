package main

import (
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	// hashAgorithm()
	// SymmetricKeyAlorithm()
	// PublicKeyInfrastructure()

}

func SignAndVerify() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	message := "안녕하세요. Go언어"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		h1,
		digest,
	)
	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(
		publicKey,
		h2,
		digest,
		signature,
	)

	if err != nil {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}

}

func PublicKeyInfrastructure() {
	s := `동해 물과 백두산이 마르고 닳도록 하느님이 보우하나 우리나라 만세.
	무궁화 삼천리 화려강산 대한사람 대한으로 길이 보전하세.`

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := &privateKey.PublicKey

	cipherText, _ := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)

	plainText, _ := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		[]byte(cipherText),
	)
	fmt.Println(s)
	fmt.Println(cipherText)
	fmt.Println(plainText)

}

func SymmetricKeyAlorithm() {
	fmt.Println("SKA")
	key := "secret-123451234"
	s := "shake!-123451234"

	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
		return
	}
	cipherText := make([]byte, cipherBlock.BlockSize())
	cipherBlock.Encrypt(cipherText, []byte(s))

	fmt.Printf("%x\n", cipherText)
}

func HashAgorithm() {
	s := "Hello, World!"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("World!"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
}
