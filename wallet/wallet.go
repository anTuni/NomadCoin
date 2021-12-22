package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/anTuni/NomadCoin/utils"
)

const PrivateKey string = "307702010104200b546c88104d5a5b3bb164c84369af98a510b2e7a5c8b2483236a44ce5e1bdb6a00a06082a8648ce3d030107a144034200048e2f263a9e7f64d1bda8f81f57d00427f783715ef09960ea2c18a60f354e387d45088a6b1348b6b2eab791885b82a7013bd54fddbcbd455779be6d3e3d2306c9"
const hashedMsg string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
const signiture string = "ee73bf4b0693a62f5095056e7b11f1a78e6f96b65fe5e7a3bbc720cd9a3f6b0aad1ede0f3b830f38c72cc3aae545ec371f67df93a8590bc0ce6a2639d94e3ed6"

func Start() {

	PrivateKeyByte, err := hex.DecodeString(PrivateKey)
	utils.HandleErr(err)
	parsedPrivate, err := x509.ParseECPrivateKey(PrivateKeyByte)
	utils.HandleErr(err)
	fmt.Printf("%x", parsedPrivate)

	sigBytes, err := hex.DecodeString(signiture)
	utils.HandleErr(err)

	var r, s = big.Int{}, big.Int{}

	Rbytes := sigBytes[:(len(sigBytes) / 2)]
	Sbytes := sigBytes[(len(sigBytes) / 2):]

	r.SetBytes(Rbytes)
	s.SetBytes(Sbytes)
	fmt.Println(r)
	fmt.Println(s)

}
