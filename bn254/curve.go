package bn254

import "github.com/sammyne/elliptic/bn254/internal/miracl/BN254"

const compressedPointLen = int(BN254.MODBYTES) + 1

var N = bigIntToGo(n)

var (
	p  = BN254.NewBIGints(BN254.Modulus)
	n  = BN254.NewBIGints(BN254.CURVE_Order)
	g1 = BN254.ECP_generator()
	g2 = BN254.ECP2_generator()
)
