package bn254

import (
	"fmt"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

const gtLen = 3 * 4 * BN254.MODBYTES

type GT struct {
	value *BN254.FP12
}

func (c *GT) Marshal() []byte {
	var out [gtLen]byte
	c.value.ToBytes(out[:])
	return out[:]
}

func (c *GT) Unmarshal(buf []byte) ([]byte, error) {
	if len(buf) < int(gtLen) {
		return nil, fmt.Errorf("buf is two small, expect at least %d", gtLen)
	}

	c.value = BN254.FP12_fromBytes(buf)
	return buf[gtLen:], nil
}

func Pair(X1 *G1, X2 *G2) *GT {
	return &GT{value: BN254.Ate(X2.value, X1.value)}
}
