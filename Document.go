package entropy

import (
	"strings";
	"regexp";
	)



type Document struct {
	text string
	word_freq map[string]uint32
}

//Function recieves sentence and breaks it up into words.
func parse_text(text string) (split []string) {
	r := regexp.MustCompile("[^a-zA-Z ]")
	text = strings.ToLower(text)
	text = r.ReplaceAllLiteralString(text,"")
	split = strings.Split(text," ")
	return
}

//TODO: Count terms for TFIF

//TODO: Divide into n-grams and loop through

//TODO: Update probs

//TODO: Test if top

//TODO: Calculate entropy

//I'll need some document-level and some overall-level datastores. 