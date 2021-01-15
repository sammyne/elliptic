package bn254

import (
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

type G2 struct {
	value *BN254.ECP2
}

func (c *G2) Add(a, b *G2) *G2 {
	cc := new(BN254.ECP2)
	cc.Copy(a.value)
	cc.Add(b.value)

	return &G2{value: cc}
}

func (c *G2) Marshal() []byte {
	var out [compressedPointLen]byte
	c.value.ToBytes(out[:], true)
	return out[:]
}

func (c *G2) Neg(a *G2) *G2 {
	aa := new(BN254.ECP2)
	aa.Copy(a.value)
	aa.Neg()

	return &G2{value: aa}
}

func (c *G2) ScalarBaseMult(k *big.Int) *G2 {
	return &G2{value: BN254.G2mul(c.value, bigIntFromGo(k))}
}

func (c *G2) ScalarMult(a *G2, k *big.Int) *G2 {
	return &G2{value: BN254.G2mul(a.value, bigIntFromGo(k))}
}

func (c *G2) Set(a *G2) *G2 {
	if c.value == nil {
		c.value = new(BN254.ECP2)
	}

	c.value.Copy(a.value)

	return c
}

func (c *G2) String() string {
	return c.value.ToString()
}

func (c *G2) Unmarshal(m []byte) ([]byte, error) {
	if len(m) < compressedPointLen {
		return nil, errors.New("too few bytes")
	}

	c.value = BN254.ECP2_fromBytes(m)

	return m[compressedPointLen:], nil
}

func NewG2(r io.Reader) (*big.Int, *G2, error) {
	x, err := randBigInt(r)
	if err != nil {
		return nil, nil, fmt.Errorf("fail to generate rand x: %w", err)
	}

	X := &G2{value: BN254.G2mul(g2, x)}

	return bigIntToGo(x), X, nil
}
