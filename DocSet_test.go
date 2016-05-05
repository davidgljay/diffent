package entropy

import (
	"testing";
	// "strings";
	"reflect";
	// "fmt";
	)

func TestNewWord(t *testing.T) {
	word := NewWord("test")
	if (reflect.TypeOf(word).String() != "*entropy.word") {
		t.Error("Error creating word: " + reflect.TypeOf(word).String())
		t.Fail()
	}

	if (reflect.TypeOf(word.prob_map).String() != "map[string]float32") {
		t.Error("Error creating probability map in word")
		t.Fail()
	}
}