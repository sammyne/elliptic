package bn254

import (
	"fmt"
	"io"
	"math/big"

	"github.com/sammyne/elliptic/bn254/internal/miracl/BN254"
)

// bigInt in little endian
type bigInt = BN254.BIG

func bigIntFromGo(v *big.Int) *bigInt {
	return BN254.FromBytes(v.Mod(v, N).Bytes())
}

func bigIntToGo(v *bigInt) *big.Int {
	var b [BN254.MODBYTES]byte
	v.ToBytes(b[:])

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
