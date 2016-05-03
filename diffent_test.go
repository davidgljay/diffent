package diffent

import (
	"testing";
//	"fmt";
	"strings";
	)

func TestParse(t *testing.T) {
	parsed := parse_text("This! is a test of stuff and things.")
	joined := strings.Join(parsed," ")
	if (joined != "this is a test of stuff and things") {
		t.Error("Incorrect parsing: " + joined)
		t.Fail()
	}
}