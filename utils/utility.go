//Package utils contains fuctions used across the application
//	utils.{function name }
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var logfn = log.Panic

//Function handling error
func HandleErr(err error) {
	if err != nil {
		logfn(err)
	}
}

//Fuction ToBytes takes a interface and encode it to slice of byte ans return that
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}
func FromBytes(i interface{}, data []byte) {

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	HandleErr(err)
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}

func ToJSON(i interface{}) []byte {
	bytes, err := json.Marshal(i)
	HandleErr(err)
	return bytes
}
