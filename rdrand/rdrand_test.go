package rdrand

import "testing"

func TestRdrand(t *testing.T) {
	t.Log(RDRand())
	t.Log(RDSeed())
}
