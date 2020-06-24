package byteflow

import (
	"bytes"
	"testing"
)

func TestTrimOperation_Execute(t *testing.T) {
	buf := []byte(" \thello\n")
	out, i, ok := Run(buf, TrimOp{Left: true})
	if i != 1 {
		t.Fatal("1")
	}
	if !ok {
		t.Fatal("2")
	}
	if string(out) != "hello\n" {
		t.Fatal("3")
	}
	out, i, ok = Run(out, TrimOp{Left: true, Right: true})
	if i != 1 {
		t.Fatal("1")
	}
	if !ok {
		t.Fatal("2")
	}
	if string(out) != "hello" {
		t.Fatal("3")
	}

}

func TestRuneOperation_Execute(t *testing.T) {
	buf := []byte{0x01, 0x02, 0x03}
	var r rune
	out, i, ok := Run(buf, RuneOp{
		Remove:  false,
		Out:     &r,
		Allowed: []rune{0x01, 0x04},
	})
	if !ok {
		t.Fatal("1")
	}
	if i != 1 {
		t.Fatal("2")
	}
	if !bytes.Equal(out, buf) {
		t.Fatal("3")
	}
	if r != 0x01 {
		t.Fatal("4")
	}
	out, i, ok = Run(buf, RuneOp{
		Remove:  true,
		Out:     &r,
		Allowed: []rune{0x01, 0x04},
	})
	if !ok {
		t.Fatal("5")
	}
	if i != 1 {
		t.Fatal("6")
	}
	if !bytes.Equal(out, []byte{0x02, 0x03}) {
		t.Fatal("7")
	}
	if r != 0x01 {
		t.Fatal("8")
	}
	out, i, ok = Run(buf, RuneOp{
		Remove: true,
		Out:    &r,
	})
	if !ok {
		t.Fatal("5")
	}
	if i != 1 {
		t.Fatal("6")
	}
	if !bytes.Equal(out, []byte{0x02, 0x03}) {
		t.Fatal("7")
	}
	if r != 0x01 {
		t.Fatal("8")
	}
}

func TestNumberOperation_Execute(t *testing.T) {
	buf := []byte("2006/08/17")
	var y, m, d int64
	out, i, ok := Run(
		buf,
		IntOp{Len: 4, Remove: true, Base: 10, Out: &y},
		RuneOp{Remove: true, Allowed: []rune{'/'}},
		IntOp{Len: 2, Remove: true, Base: 10, Out: &m},
		RuneOp{Remove: true, Allowed: []rune{'/'}},
		IntOp{Len: 2, Remove: true, Base: 10, Out: &d},
	)
	if !ok {
		t.Fatal("1", i)
	}
	if len(out) != 0 {
		t.Fatal("2")
	}
	if i != 5 {
		t.Fatal("3")
	}
	if y != 2006 {
		t.Fatal("4")
	}
	if m != 8 {
		t.Fatal("5")
	}
	if d != 17 {
		t.Fatal("6")
	}
}

func TestMarkDecodeOperation_Execute(t *testing.T) {
	buf := []byte("K[v1]hello K[v2,v3]K[v4]\tK[v5] K[v6")
	var k string
	out, i, ok := Run(buf, MarkDecodeOp{Name: "K", Out: &k, Combine: true, Separator: ","})
	if !bytes.Equal(out, buf) {
		t.Fatal("1")
	}
	if i != 1 {
		t.Fatal("2")
	}
	if !ok {
		t.Fatal("3")
	}
	if k != "v1,v2,v3,v5" {
		t.Fatal("4", k)
	}
	k = ""
	buf = []byte("K[v1]hello K[v2,v3]K[v4]\tK[v5] K[] ")
	out, i, ok = Run(buf, MarkDecodeOp{Name: "K", Out: &k, Combine: true, Separator: ","})
	if !bytes.Equal(out, buf) {
		t.Fatal("1")
	}
	if i != 1 {
		t.Fatal("2")
	}
	if !ok {
		t.Fatal("3")
	}
	if k != "v1,v2,v3,v5" {
		t.Fatal("4", k)
	}
}

func TestJSONDecodeOperation_Execute(t *testing.T) {
	buf := []byte(`{"hello":"world"}`)
	var m map[string]interface{}
	out, i, ok := Run(buf, JSONDecodeOp{Remove: true, Out: &m})
	if i != 1 {
		t.Fatal("1")
	}
	if !ok {
		t.Fatal("2")
	}
	if len(out) != 0 {
		t.Fatal("3")
	}
	if m["hello"].(string) != "world" {
		t.Fatal("4")
	}
}
