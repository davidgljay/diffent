package entropy

import "container/list"

//TODO: Need to figure out how to go make lists part of a structs
//Do lists work for docs and mentions? 
//Probably, I generally will be iterating through all mentions in an export
//I can always convert to an array for analysis

type DocSet struct {
	documents list
	words map[string]word
	ngram int
}

type word struct {
	freq uint32
	prob_array map[string]float32
	mentions list
}

type mention struct {
	ts uint64
	entropy float32
	top []probability
}

type probability struct {
	word string
	prob float32
}

func NewDocSet(ngram int) (docset *DocSet) {
	docset.words = make(map[string]word)
	docset.ngram = ngram
	return
}

func NewWord(w string) (wrd *word) {
	wrd.prob_array = make(map[string]float32)
	return wrd
}

func AddDoc(self *DocSet, doc string) {
	self.documents.PushBack(doc)
}


//TODO: Iterate through docs again and get entropy calc

//TODO: Output to file

