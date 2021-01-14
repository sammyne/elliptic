package bn254

import (
	"fmt"
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

// bigInt in little endian
type bigInt = BN254.BIG

func bigIntFromBE(v *big.Int) *bigInt {
	panic("todo")
}

func bigIntToBE(v *bigInt) *big.Int {
	var b [BN254.MODBYTES]byte
	v.ToBytes(b[:])

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return new(big.Int).SetBytes(b[:])
}

func randBigInt(r io.Reader) (*bigInt, error) {
	var b [BN254.MODBYTES]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return nil, fmt.Errorf("lacking entropy: %w", err)
	}

	v := BN254.FromBytes(b[:])
	v.Mod(n)

	return v, nil
}
