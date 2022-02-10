package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	str := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"

	t.Run("Hash is always same", func(t *testing.T) {
		s := struct{ Test string }{Test: "test"}
		hash := Hash(s)
		if hash != str {
			t.Error("Hash is not same all the time")
		}
	})

	t.Run("Hash is hex encoded", func(t *testing.T) {
		s := struct{ Test string }{Test: "test"}
		hash := Hash(s)
		_, err := hex.DecodeString(hash)
		if err != nil {
			t.Error("Hash isn't hex encoded")
		}
	})
}

func ExampleHash() {
	s := struct{ Test string }{Test: "test"}
	hash := Hash(s)
	fmt.Println(hash)
	//Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}
