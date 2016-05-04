package entropy

import (
	"testing";
	"strings";
	"reflect";
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