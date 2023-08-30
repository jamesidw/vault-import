package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/tink-crypto/tink-go/v2/kwp/subtle"
	"os"
)

func main() {
	wrappingKeyString, err := os.ReadFile("./wrapping.pem")
	if err != nil {
		panic(err)
	}
	targetKeyString, err := os.ReadFile("./target.pem")
	if err != nil {
		panic(err)
	}

	ephemeralAESKey := make([]byte, 32)

	keyBlock, _ := pem.Decode(wrappingKeyString)
	parsedKey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	targetBlock, _ := pem.Decode(targetKeyString)

	wrapKWP, err := subtle.NewKWP(ephemeralAESKey)
	if err != nil {
		panic(err)
	}
	wrappedTargetKey, err := wrapKWP.Wrap(targetBlock.Bytes)
	if err != nil {
		panic(err)
	}
	wrappedAESKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		parsedKey.(*rsa.PublicKey),
		ephemeralAESKey,
		[]byte{},
	)
	if err != nil {
		panic(err)
	}
	combinedCiphertext := append(wrappedAESKey, wrappedTargetKey...)
	base64Ciphertext := base64.StdEncoding.EncodeToString(combinedCiphertext)
	fmt.Print(base64Ciphertext)
}
