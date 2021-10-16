package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// fmt.printfをテストする、依存関係を持っている関数をテスト
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	if got != want {
		t.Errorf("want: %q but got: %q", want, got)
	}
}

func TestPrint(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	print("hoge")
	fmt.Println("hoge")
	print("hoge")
	print("hoge")

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	fmt.Println(buf)

	if buf.String() != "hoge" {
		t.Errorf("print() = %s, want hoge", buf.String())
	}
}
