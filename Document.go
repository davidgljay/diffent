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
func updateProbs (self *Document, docSet *DocSet) {

	//Start by counting through the first ngram of words and updating prob_maps
	for i := 1; i < docSet.ngram-1; i++ {
		iterateNgramAndUpdate(self, docSet, i, i+1)
	}

	//Next, move through each additional word, updating the prob_map for the words behind it
	for i := docSet.ngram-1; i < len(self.words); i++ {
		iterateNgramAndUpdate(self, docSet, i, docSet.ngram)
	}
}

func iterateNgramAndUpdate(self * Document, docSet *DocSet, i int, jmax int) {
		word := self.words[i]
		//Calculate the TFIF
		wordTFIF := tfif(self.word_freq[word], docSet.words[word].freq)
		//Count backwards through ngram
		for j := 1; j < jmax; j++ {
			compword := self.words[i-j]
			compTFIF := tfif(self.word_freq[word], docSet.words[word].freq)

			//Update the probability with tfif
			docSet.words[word].prob_map[compword] += compTFIF
			docSet.words[compword].prob_map[word] += wordTFIF
		}
}

func tfif (docCount uint32, setCount uint32) (tfif float64) {
	tfif = float64(docCount)/float64(setCount)
	return
}
