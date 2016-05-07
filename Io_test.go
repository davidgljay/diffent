package entropy

import (
	"testing";
	// "fmt";
	)

func TestWriteToFile(t *testing.T) {
	WriteToFile("test.txt", "Test test test!")
}