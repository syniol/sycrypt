// Copyright 2023 Syniol Limited.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// Package sycrypt provides a simple way to create a military grade passwords

//	package main
//
//	import (
//		"encoding/json"
//
//		"github.com/syniol/sycrypt"
//	)
//
//	func main() {
//		credentials, err := sycrypt.NewCredential("johnspassword1")
//		if err != nil {
//			panic(err)
//		}
//
//		// prints hashed password
//		result, err := json.Marshal(credentials)
//		if err != nil {
//			panic(err)
//		}
//
//		println(string(result))
//
//		// prints verification status true
//		isVerified := credentials.VerifyPassword("johnspassword1")
//		println("verification status", isVerified)
//
//		// prints verification status false
//		isVerified = credentials.VerifyPassword("alicepassword1")
//		println("verification status", isVerified)
//	}

package sycrypt

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

// Credential is a Data Structure where values for both private & public keys are stored with Hashed password
type Credential struct {
	// Key (public key) is represented in base64 string before hex encode
	Key string `json:"key"`

	// HashedPassword is represented in base64 string before hex encode (This is hashed password using PublicKey & PrivateKey)
	HashedPassword string `json:"hashedPassword"`
}

// NewCredential will create a new Credential structure with given password as a parameter
func NewCredential(password string) (*Credential, error) {
	public, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal public key: %s", err.Error())
	}

	var pemEncodedPublicKey bytes.Buffer
	if err := pem.Encode(&pemEncodedPublicKey, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}); err != nil {
		return nil, err
	}

	return &Credential{
		Key: encodeHash(pemEncodedPublicKey.Bytes()),
		HashedPassword: encodeHash(
			ed25519.Sign(private, []byte(password)),
		),
	}, nil
}

// VerifyPassword will verify plain password given as a string parameter inputPassword
func (cred *Credential) VerifyPassword(inputPassword string) bool {
	return ed25519.Verify(
		decodePublicCert(decodeHash([]byte(cred.Key))),
		[]byte(inputPassword),
		decodeHash([]byte(cred.HashedPassword)),
	)
}

func decodeHash(hash []byte) []byte {
	base64Hash := make([]byte, base64.StdEncoding.DecodedLen(len(hash)))
	base64.StdEncoding.Decode(base64Hash, hash)

	hexHash := make([]byte, hex.DecodedLen(len(base64Hash)))
	hex.Decode(hexHash, base64Hash)

	return hexHash
}

func encodeHash(plaintext []byte) string {
	hexHash := make([]byte, hex.EncodedLen(len(plaintext)))
	hex.Encode(hexHash, plaintext)

	base64Hash := make([]byte, base64.StdEncoding.EncodedLen(len(hexHash)))
	base64.StdEncoding.Encode(base64Hash, hexHash)

	return string(base64Hash)
}

func decodePublicCert(cert []byte) []byte {
	out, _ := pem.Decode(cert)

	parsedPublicKey, _ := x509.ParsePKIXPublicKey(out.Bytes)

	return parsedPublicKey.(ed25519.PublicKey)
}
