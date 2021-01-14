package bn254

import (
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

type G2 struct {
	value BN254.ECP2
}

func (c *G2) Add(a, b *G2) *G2 {
	panic("todo")
}

func (c *G2) Marshal() []byte {
	panic("todo")
}

func (c *G2) Neg(a *G2) *G2 {
	panic("todo")
}

func (c *G2) ScalarBaseMult(k *big.Int) *G2 {
	panic("todo")
}

func (c *G2) ScalarMult(a *G2, k *big.Int) *G2 {
	panic("todo")
}

func (c *G2) String() string {
	panic("todo")
}

func (c *G2) Unmarshal(m []byte) ([]byte, error) {
	panic("todo")
}

func NewG2(r io.Reader) (*big.Int, *G2, error) {
	panic("todo")
}
