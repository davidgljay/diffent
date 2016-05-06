package entropy

import (
	"container/list"
	"sort"
	// "fmt"
	)

//TODO: Need to figure out how to go make lists part of a structs
//Do lists work for docs and mentions? 
//Probably, I generally will be iterating through all mentions in an export
//I can always convert to an array for analysis

type DocSet struct {
	documents list.List
	words map[string]*word
	ngram int
}

type word struct {
	freq uint32
	prob_map map[string]float32
	mentions list.List
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

type ByProb []probability

func (a ByProb) Len() int           { return len(a) }
func (a ByProb) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByProb) Less(i, j int) bool { return a[i].prob > a[j].prob }


func NewDocSet(ngram int) (docset *DocSet) {
	docset = new(DocSet)
	docset.words = make(map[string]*word)
	docset.ngram = ngram
	return
}

func NewWord(w string) (wrd *word) {
	wrd = new(word)
	wrd.prob_map = make(map[string]float32)
	return
}

func AddDoc(self *DocSet, doc string) {
	self.documents.PushBack(doc)
}

func TestIfTop(self *DocSet, word1 string, word2 string, top []probability) ([]probability) {
	//Right now just tests a words match with itself, need to test across all words
	p := self.words[word1].prob_map[word2]
	if (p > top[9].prob) {
			top[10] = probability{word2, p}
			sort.Sort(ByProb(top))
	}
	return top
}


//TODO: Iterate through docs again and get entropy calc

//TODO: Output to file

//TODO: Test if top

//TODO: Calculate entropy
