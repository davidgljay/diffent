package entropy

import (
	"testing";
	"strings";
	"reflect";
	// "fmt";
	)

func TestParse(t *testing.T) {
	parsed := parseText("This! is a test of stuff and things.")
	joined := strings.Join(parsed," ")
	if (joined != "this is a test of stuff and things") {
		t.Error("Incorrect parsing: " + joined)
		t.Fail()
	}
}

func TestNewDocument(t *testing.T) {
	document := NewDocument("for there is no folly of the beast of the earth which is not infinitely outdone by the madness of men")
	if (len(document.words) != 21) {
		t.Error("Error parsing text in document creation.")
		t.Fail()
	}
	if (reflect.TypeOf(document.word_freq).String() != "map[string]uint32") {
		t.Error("Document failed to initialize map")
		t.Fail()
	}
}

func TestCountTerms(t *testing.T) {
	document := NewDocument("whales whales whales!")
	countTerms(document)
	if (document.word_freq["whales"] != 3) {
		t.Error("Error in count terms function")
		t.Fail()
	}
}

func TestUpdateProbs(t *testing.T) {
	//Set up DocSet
	docSet := NewDocSet(3)

	//Set up Doc
	doc := NewDocument("Whales whales whales are great!")
	countTerms(doc)
	for i := 0; i < len(doc.words); i++ {
		docSet.words[doc.words[i]] = NewWord()
		docSet.words[doc.words[i]].freq = 10
	}
	if (tfif(doc.word_freq["whales"],docSet.words["whales"].freq) != 0.3)  {
		t.Error("Incorrect calculation of TFIF in update")
		t.Fail()
	}

	//Test
	updateProbs(doc, docSet)
	if (docSet.words["whales"].prob_map["whales"] != 1.8) {
		t.Error("Failure counting terms in updateProbs, should be 1.8")
		t.Error(docSet.words["whales"].prob_map["whales"])
		t.Fail()
	}
}