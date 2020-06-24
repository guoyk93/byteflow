package byteflow

import (
	"bytes"
	"testing"
)

type trueOperation int

func (trueOperation) Execute(in []byte) (out []byte, ok bool) {
	return in, true
}

type falseOperation int

func (falseOperation) Execute(in []byte) (out []byte, ok bool) {
	return in, false
}

func TestRun(t *testing.T) {
	buf := []byte{0x01, 0x02, 0x03}
	out, i, ok := Run(buf, trueOperation(0), trueOperation(0), falseOperation(0), trueOperation(0))
	if !bytes.Equal(out, buf) {
		t.Fatal("1")
	}
	if i != 2 {
		t.Fatal("2")
	}
	if ok {
		t.Fatal("3")
	}
}
