package main

import (
	"fmt"
	"log"

	encrypt "github.com/nodias/go_begin/exercise/encryption/encrypt"
)

func main() {
	plainText := "hello, my boy"

	privateKey, publicKey, err := encrypt.RsaKeyGenerator()
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := encrypt.RsaKeyEncode(publicKey, plainText)
	if err != nil {
		log.Fatal(err)
	}

	if plainText2, err := encrypt.RsaKeyDecode(*privateKey, ciphertext); err != nil {
		log.Fatal(err)
	} else if plainText2 != plainText {
		fmt.Println("wrong")
		return
	}
	fmt.Println("login success")

	message := "hello, son. how are you?"
	digest := encrypt.ToHash(message)
	signature, err := encrypt.SingMessage(privateKey, digest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message sent succesfully")

	err = encrypt.VertifyMessage(publicKey, signature, digest, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receive Message is : ", message)
}
