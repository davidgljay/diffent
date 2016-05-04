package entropy

import (
	"strings";
	"regexp";
	)



type Document struct {
	words []string
	word_freq map[string]uint32
}

//Initialize a document
func NewDocument (text string) *Document {
	document := new(Document)
	document.words = parseText(text)
	document.word_freq = make(map[string]uint32)
	return document
}

//Function recieves sentence and breaks it up into words.
func parseText(text string) (split []string) {
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