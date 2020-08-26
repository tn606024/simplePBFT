package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var privateKey8080 		*rsa.PrivateKey
var privateKey8081 		*rsa.PrivateKey
var privateKey8082 		*rsa.PrivateKey
var privateKey8083 		*rsa.PrivateKey
var privateKeyClient  	*rsa.PrivateKey
var publicKey8080  		*rsa.PublicKey
var publicKey8081  		*rsa.PublicKey
var publicKey8082  		*rsa.PublicKey
var publicKey8083  		*rsa.PublicKey
var publicKeyClient 	*rsa.PublicKey
var KnownNodes []*KnownNode
var KeypairMap map[int]Keypair
var ClientNode *KnownNode
func init(){
	var err error
	generateKeyFiles()
	privateKey8080, publicKey8080, err = getKeyPairByFile(0)
	if err != nil {
		panic(err)
	}
	privateKey8081, publicKey8081, err = getKeyPairByFile(1)
	if err != nil {
		panic(err)
	}
	privateKey8082, publicKey8082, err = getKeyPairByFile(2)
	if err != nil {
		panic(err)
	}
	privateKey8083, publicKey8083, err = getKeyPairByFile(3)
	if err != nil {
		panic(err)
	}
	privateKeyClient, publicKeyClient, err = getKeyPairByFile(4)
	if err != nil {
		panic(err)
	}
	KnownNodes = []*KnownNode{
		{
			0,
			"localhost:8080",
			publicKey8080,
		},
		{
			1,
			"localhost:8081",
			publicKey8081,
		},
		{
			2,
			"localhost:8082",
			publicKey8082,
		},
		{
			3,
			"localhost:8083",
			publicKey8083,
		},
	}
	KeypairMap = map[int]Keypair{
		0:{
			privateKey8080,
			publicKey8080,
		},
		1:{
			privateKey8081,
			publicKey8081,
		},
		2:{
			privateKey8082,
			publicKey8082,
		},
		3:{
			privateKey8083,
			publicKey8083,
		},
		4:{
			privateKeyClient,
			publicKeyClient,
		},
	}
	ClientNode = &KnownNode{
		4,
		"localhost:8084",
		publicKeyClient,
	}
}

func getKeyPairByFile(nodeID int) (*rsa.PrivateKey, *rsa.PublicKey, error){
	privFile, _ := filepath.Abs(fmt.Sprintf("./Keys/%d_priv",nodeID))
	pubFile, _ := filepath.Abs(fmt.Sprintf("./Keys/%d_pub",nodeID))
	fbytes, err := ioutil.ReadFile(privFile)
	if err != nil {
		return nil,nil, err
	}
	block, _:= pem.Decode(fbytes)
	if block == nil {
		return nil,nil, fmt.Errorf("parse block occured error")
	}
	privkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil{
		return nil,nil, err
	}
	pubfbytes, err := ioutil.ReadFile(pubFile)
	if err != nil {
		return nil,nil, err
	}
	pubblock, _:= pem.Decode(pubfbytes)
	if pubblock == nil {
		return nil,nil, fmt.Errorf("parse block occured error")
	}
	pubkey, err := x509.ParsePKIXPublicKey(pubblock.Bytes)
	if err != nil{
		return nil,nil, err
	}
	return privkey, pubkey.(*rsa.PublicKey), nil
}

func generateKeyFiles(){
	if !FileExists("./Keys") {
		err := os.Mkdir("Keys", 0700)
		if err != nil {
			panic(err)
		}
		for i := 0; i<=4; i++ {
			filename, _ := filepath.Abs(fmt.Sprintf("./Keys/%d",i))
			if !FileExists(filename + "_priv") && !FileExists(filename + "_pub"){
				priv, pub := generateKeyPair()
				err := ioutil.WriteFile(filename+"_priv", priv, 0644)
				if err != nil {
					panic(err)
				}
				ioutil.WriteFile(filename+"_pub", pub, 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func generateKeyPair() ([]byte, []byte){
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	mprivkey := x509.MarshalPKCS1PrivateKey(privkey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: mprivkey,
	}
	bprivkey := pem.EncodeToMemory(block)
	pubkey := &privkey.PublicKey
	mpubkey, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: mpubkey,
	}
	bpubkey := pem.EncodeToMemory(block)
	return bprivkey, bpubkey
}