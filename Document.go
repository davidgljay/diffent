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

//Count terms for TFIF
func countTerms (self *Document) {
	for i := 0; i < len(self.words); i++ {
		self.word_freq[self.words[i]]++
	}
}

//Divide into n-grams and loop through
func updateProbs (self *Document, docSet DocSet) {
	for i := docSet.ngram-1; i < len(self.words); i++ {
		word := self.words[i]

		//Calculate the TFIF
		wordTFIF := tfif(self.word_freq[word], docSet.words[word].freq)
		Count backwards through ngram
		for j := 1; j < docSet.ngram; j++ {
			compword := self.words[i-j]
			compTFIF := tfif(self.word_freq[word], docSet.words[word].freq)

			//Update the probability with te 
			docSet.words[word].prob_array[compword] += compTFIF
			docSet.words[compword].prob_array[word] += wordTFIF
		}
	}
}

func tfif (docCount int32, setCount int32) (tfif float32) {
	tfif = float32(docCount)/float32(setCount)
	return
}


//TODO: Update probs

//TODO: Test if top

//TODO: Calculate entropy

//I'll need some document-level and some overall-level datastores. 