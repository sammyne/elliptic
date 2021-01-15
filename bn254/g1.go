package bn254

import (
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

type G1 struct {
	value *BN254.ECP
}

func (c *G1) Add(a, b *G1) *G1 {
	cc := new(BN254.ECP)
	cc.Copy(a.value)
	cc.Add(b.value)

	return &G1{value: cc}
}

func (c *G1) Marshal() []byte {
	var out [compressedPointLen]byte
	c.value.ToBytes(out[:], true)
	return out[:]
}

func (c *G1) Neg(a *G1) *G1 {
	aa := new(BN254.ECP)
	aa.Copy(a.value)
	aa.Neg()

	return &G1{value: aa}
}

func (c *G1) ScalarBaseMult(k *big.Int) *G1 {
	return &G1{value: BN254.G1mul(c.value, bigIntFromGo(k))}
}

func (c *G1) ScalarMult(a *G1, k *big.Int) *G1 {
	return &G1{value: BN254.G1mul(a.value, bigIntFromGo(k))}
}

func (c *G1) Set(a *G1) *G1 {
	if c.value == nil {
		c.value = new(BN254.ECP)
	}

	c.value.Copy(a.value)

	return c
}

func (c *G1) String() string {
	return c.value.ToString()
}

func (c *G1) Unmarshal(m []byte) ([]byte, error) {
	if len(m) < compressedPointLen {
		return nil, errors.New("too few bytes")
	}

	c.value = BN254.ECP_fromBytes(m)

	return m[compressedPointLen:], nil
}

func NewG1(r io.Reader) (*big.Int, *G1, error) {
	x, err := randBigInt(r)
	if err != nil {
		return nil, nil, fmt.Errorf("fail to generate rand x: %w", err)
	}

	X := &G1{value: BN254.G1mul(g1, x)}

	return bigIntToGo(x), X, nil
}
