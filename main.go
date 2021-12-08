package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficulty := 5
	target := strings.Repeat("0", difficulty)
	nonce := 1
	for {
		hash := sha256.Sum256([]byte("string" + fmt.Sprint(nonce)))
		fmt.Printf("Hash : %x\nTarget : %s\nNonce : %d\n", hash, target, nonce)
		if strings.HasPrefix(fmt.Sprintf("%x", hash), target) {
			return
		} else {
			nonce++
		}

	}
}
