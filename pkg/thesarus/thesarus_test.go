package thesarus

import (
	"os"
	"testing"
)

func TestSynonym(t *testing.T) {
	term := "code"
	bht := New(os.Getenv("BHT_KEY"))

	syns, err := bht.Synonyms(term)

	if err != nil {
		t.Fatal(err)
	}

	if len(syns) == 0 {
		t.Fatalf("expected synonyms. found empyt slice")
	}
}
