package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Hola mundo!\n")
	_, got := count(b, false, false)
	want := 2
	if got != want {
		t.Errorf("Expected %d, got %d instead\n", want, got)
	}

}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	got := 3
	_, want := count(b, true, false)
	if got != want {
		t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Hola mundo!😊\n")
	want := 16
	_, got := count(b, false, true)
	if want != got {
		t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}
