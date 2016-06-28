package btext

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Parsetest
func TestParsetest(t *testing.T) {
	text := BParseLine("1 2 3 4;fa")
	if !reflect.DeepEqual(text, []byte{1, 2, 3, 4}) {
		t.Error("not same", text)
	}
}

func TestParseFile(t *testing.T) {
	binarys := BParseFile("bin.txt")
	test := []byte{0, 1, 10, 16}
	if !reflect.DeepEqual(test, binarys) {
		t.Error("not same", binarys)
	}
}

func TestParseLog(t *testing.T) {
	binary := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 0x20, 1, 2, 3, 4, 5, 6, 7, 8}
	rst := TParseAry(binary)
	data, err := ioutil.ReadFile("rst.txt")
	if err != nil {
		t.Error("can`t read")
	}
	text := strings.TrimRight(string(data), "\n")
	if rst != text {
		t.Error("not same\n", text, "\n", rst, "")
	}
}
