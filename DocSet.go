package entropy

type DocSet struct {
	documents []string
	words []word
}


type word struct {
	freq uint32
	prob_array []probability
	mentions []mention
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
//TODO: Create struct

//TODO: Iterate through docs for TFIF

//TODO: Iterate through docs again and get entropy calc

//TODO: Output to file

