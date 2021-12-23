package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/anTuni/NomadCoin/utils"
)

const (
	fileName = "nomadcoinWallet"
)

type wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}
func createPrivKey() *ecdsa.PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return key
}
func persistPrivKey(k *ecdsa.PrivateKey) {
	marshaled, err := x509.MarshalECPrivateKey(k)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, marshaled, 0644)
	utils.HandleErr(err)
}
func restorePrivateKey() *ecdsa.PrivateKey {
	keyAsByte, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err := x509.ParseECPrivateKey(keyAsByte)
	utils.HandleErr(err)
	return key
}
func AfromK(key *ecdsa.PrivateKey) string {
	bytes := append(key.X.Bytes(), key.Y.Bytes()...)
	return fmt.Sprintf("%x", bytes)
}
func sign(payload string, w *wallet) string {
	hashAsByte, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hashAsByte)
	utils.HandleErr(err)
	z := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", z)
}
func Verify(signiture, payload, publicKey string) {
}
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.PrivateKey = restorePrivateKey()
		} else {
			privKey := createPrivKey()
			w.PrivateKey = privKey
			persistPrivKey(privKey)
		}
		w.Address = AfromK(w.PrivateKey)
	}
	return w
}
