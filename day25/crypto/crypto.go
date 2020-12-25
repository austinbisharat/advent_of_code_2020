package crypto

import "math/big"

func Exp(base, power, mod int64) int64 {
	return new(big.Int).Exp(big.NewInt(base), big.NewInt(power), big.NewInt(mod)).Int64()
}

func DiscreteLog(base, power, mod int64) int64 {
	for i := int64(2); i < mod; i++ {
		if Exp(base, i, mod) == power {
			return i
		}
	}
	panic("should never happen")
}
