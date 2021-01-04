package main

import (
	"testing"
)

func TestCountWordsNormal(t *testing.T) {
	ch := make(chan map[string]int)

	dict1 := CountWordsNormal("./texts/text1")
	go CountWordsConcurrently("./texts/text1", ch)

	dict2 := <-ch
	exp1 := dict1["100"]
	got1 := dict2["100"]

	if exp1 != got1 {
		t.Fatal("failed")
	}

	exp2 := len(dict1)
	got2 := len(dict2)

	if exp2 != got2 {
		t.Fatal("failed")
	}
}

func BenchmarkCountWordsNormal(b *testing.B) {

	for n := 0; n < b.N; n++ {
		_ = CountWordsNormal("./texts/text1")

	}
}

func BenchmarkCountWordsConcurrently(b *testing.B) {
	ch := make(chan map[string]int)
	for n := 0; n < b.N; n++ {
		go CountWordsConcurrently("./texts/text1", ch)

	}
}
