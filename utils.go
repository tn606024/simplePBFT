package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func generateDigest(msg interface{}) []byte{
	bmsg, _ := json.Marshal(msg)
	hash := sha256.Sum256(bmsg)
	return hash[:]
}

func signMessage(msg interface{}, privkey *rsa.PrivateKey) ([]byte, error){
	dig := generateDigest(msg)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, dig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func verifyDigest(msg interface{}, digest string) bool{
	return hex.EncodeToString(generateDigest(msg)) == digest
}

func verifySignatrue(msg interface{}, sig []byte, pubkey *rsa.PublicKey) (bool, error){
	dig := generateDigest(msg)
	err := rsa.VerifyPKCS1v15(pubkey,crypto.SHA256, dig, sig)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FileExists(filename string) bool {
	path, _ := filepath.Abs(filename)
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}