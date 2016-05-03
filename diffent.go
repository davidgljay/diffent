package diffent

import (
	"strings";
	"regexp";
	)

//Function recieves sentence and breaks it up into words.
func parse_text(text string) (split []string) {
	r := regexp.MustCompile("[^a-zA-Z ]")
	text = strings.ToLower(text)
	text = r.ReplaceAllLiteralString(text,"")
	split = strings.Split(text," ")
	return
}