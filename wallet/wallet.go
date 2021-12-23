package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/anTuni/NomadCoin/utils"
)

const (
	fileName = "nomadcoinWallet"
)

type wallet struct {
	PrivateKey *ecdsa.PrivateKey
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
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			//yes restore Private key from the file
		} else {
			privKey := createPrivKey()
			w.PrivateKey = privKey
			persistPrivKey(privKey)
		}
	}
	return w
}
