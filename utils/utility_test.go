package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	t.Logf("Return of function ToByte is a %s", k)
	if k != reflect.Slice {
		t.Errorf("Return of function ToByte is not a Slice then is a %s", k)
	}
}

func TestSplitter(t *testing.T) {
	type test struct {
		s      string
		sep    string
		i      int
		output string
	}
	var tests = []test{
		{s: "0:7:0", sep: ":", i: 1, output: "7"},
		{s: "070", sep: ":", i: 1, output: ""},
		{s: "0:7:0", sep: "/", i: 1, output: ""},
	}
	for _, test := range tests {
		out := Splitter(test.s, test.sep, test.i)
		if out != test.output {
			t.Errorf("Expected value : %s, function return : %s", test.output, out)
		}
	}
}

func TestHandleErr(t *testing.T) {
	var lofnBackUp = logfn
	defer func() {
		logfn = lofnBackUp
	}()

	err := errors.New("test")
	called := false
	logfn = func(v ...interface{}) {
		called = true
	}
	HandleErr(err)
	if !called {
		t.Error("HandleErr not call log.Panic()")
	}
}
func TestFromBytes(t *testing.T) {
	type testStruct struct{ Test string }
	var restore testStruct
	s := testStruct{Test: "test"}
	b := ToBytes(s)

	FromBytes(&restore, b)

	if !reflect.DeepEqual(restore, s) {
		t.Errorf("Expected %v, we got %v", s, restore)
	}
}
func TestToJSON(t *testing.T) {
	type testStruct struct{ Test string }
	s := testStruct{Test: "test"}
	var restore testStruct
	testjson := ToJSON(s)
	k := reflect.TypeOf(testjson).Kind()
	if k != reflect.Slice {
		t.Errorf("Expected : %v got : %v", reflect.Slice, k)
	}
	json.Unmarshal(testjson, &restore)
	if !reflect.DeepEqual(s, restore) {
		t.Errorf("Expected %v, we got %v", s, restore)

	}
}
