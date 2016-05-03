package diffent

import (
	"fmt";
	"strings";
	"regexp";
	)

//Function recieves sentence and breaks it up into words.
func parse_phrase(phrase string) (split []string) {
	r := regexp.MustCompile("[^a-zA-Z ]")
	phrase = strings.ToLower(phrase)
	phrase = r.ReplaceAllLiteralString(phrase,"")
	split = strings.Split(phrase," ")
	return
}