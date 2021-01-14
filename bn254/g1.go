package bn254

import (
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

type G1 struct {
	value BN254.ECP
}

func (c *G1) Add(a, b *G1) *G1 {
	panic("todo")
}

func (c *G1) Marshal() []byte {
	panic("todo")
}

func (c *G1) Neg(a *G1) *G1 {
	panic("todo")
}

func (c *G1) ScalarBaseMult(k *big.Int) *G1 {
	panic("todo")
}

func (c *G1) ScalarMult(a *G1, k *big.Int) *G1 {
	panic("todo")
}

func (c *G1) String() string {
	panic("todo")
}

func (c *G1) Unmarshal(m []byte) ([]byte, error) {
	panic("todo")
}

func NewG1(r io.Reader) (*big.Int, *G1, error) {
	panic("todo")
}
