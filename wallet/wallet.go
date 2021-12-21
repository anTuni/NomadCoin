package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/anTuni/NomadCoin/utils"
)

func Start() {

	PrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	message := "I love you"
	hashedMsg := utils.Hash(message)
	hashAsByte, err := hex.DecodeString(hashedMsg)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, PrivateKey, hashAsByte)
	utils.HandleErr(err)
	fmt.Printf("r : %d \ns : %d", r, s)
}
