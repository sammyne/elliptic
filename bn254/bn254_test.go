package bn254_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/sammyne/elliptic/bn254"
)

func TestPair(t *testing.T) {
	x1, X1, err := bn254.NewG1(rand.Reader)
	if err != nil {
		t.Fatalf("fail to rand G1: %v", err)
	}

	x2, X2, err := bn254.NewG2(rand.Reader)
	if err != nil {
		t.Fatalf("fail to rand G2: %v", err)
	}

	Z := bn254.Pair(X1, X2).Marshal()

	Y1 := new(bn254.G1).ScalarBaseMult(x2)
	Y2 := new(bn254.G2).ScalarBaseMult(x1)

	ZZ := bn254.Pair(Y1, Y2).Marshal()

	if !bytes.Equal(Z, ZZ) {
		t.Fatalf("expect %x, got %x", Z, ZZ)
	}
}
