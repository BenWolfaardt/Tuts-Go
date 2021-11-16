package mascot_test

import (
	"testing"
	"example.com/0_Official_tut/mascot"
)


func TestPack(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}