package ext

import (
	"log"
	"testing"
)

func TestExt(t *testing.T) {
	ints := Strings2Ints([]string{"123", "456", "789"})
	log.Print(ints)
	log.Print(Int64s2Strings(ints))
}
