package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//rsa 密钥文件产生
	fmt.Println(GenRsaKey(1024))
}

//GenRsaKey RSA公钥私钥产生
func GenRsaKey(bits int) (privKey, pubKey string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return "", "", err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return "", "", err
	}
	b, err := ioutil.ReadFile("private.pem")
	if err != nil {
		fmt.Print(err)
	}
	privKey = string(b)
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return "", "", err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return "", "", err
	}
	b, err = ioutil.ReadFile("public.pem")
	if err != nil {
		fmt.Print(err)
	}
	pubKey = string(b)
	defer func() {
		os.Remove("public.pem")
		os.Remove("private.pem")
	}()
	return privKey, pubKey, nil
}
