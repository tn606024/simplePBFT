package main

import "crypto/rsa"

type KnownNode struct {
	nodeID	int
	url		string
	pubkey	*rsa.PublicKey
}
