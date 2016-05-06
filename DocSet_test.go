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

func TestTestIfTop(t * testing.T) {
	//Write test for Top 10 (or another var?) function
	// testTop := make([]probability,11)
	testTop := []probability {probability {"moby",0},probability {"dick",1},probability {"was",2},probability {"a",3},probability {"happy",4},probability {"whale",5},probability {"every",6},probability {"day",7},probability {"he",8},probability {"danced",9},probability {"contra",0}}
	docSet := NewDocSet(3)
	docSet.words["white"] = NewWord("white")
	docSet.words["white"].prob_map["whales"] = 11

	newTop := TestIfTop(docSet, "white", "whales", testTop)
	if (newTop[0].word != "whales") {
		t.Error("TestTop function did not insert and sort correctly")
		t.Fail()
	}

	docSet.words["nothing"] = NewWord("nothing")
	docSet.words["nothing"].prob_map["nada"] = -0.5

	sameTop := TestIfTop(docSet, "nothing", "nada", testTop)
	if (sameTop[10].word != "moby") {
		t.Error("TestTop function failed to skip over low score")
		t.Fail()
	} 
}