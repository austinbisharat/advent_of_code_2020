package modulo

import "fmt"

// returns the unknown x in ax = b Mod base
func SolveLinearCongruence(a, b, base int) (int, int, error) {
	gcd := GCD(a, base)
	if b%gcd != 0 {
		return 0, 0, fmt.Errorf("no solution to linear congruence %d * x = %d Mod %d", a, b, base)
	}

	a /= gcd
	b /= gcd
	base /= gcd

	// ax = b mode base
	inverse := modInverse(a, base)
	return Mod(b*inverse, base), gcd, nil
}

// greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	return ExtendedGCD(a, b).GCD
}

type ExtendedGCDResult struct {
	GCD int
	X   int
	Y   int
}

func ExtendedGCD(a, b int) ExtendedGCDResult {
	prevX, x := 1, 0
	prevY, y := 0, 1
	for b > 0 {
		q := a / b
		x, prevX = prevX-q*x, x
		y, prevY = prevY-q*y, y
		a, b = b, Mod(a, b)
	}
	return ExtendedGCDResult{
		GCD: a,
		X:   prevX,
		Y:   prevY,
	}
}

// assumes gcd(a, m) = 1
func modInverse(a, m int) int {
	inv := ExtendedGCD(a, m).X
	if inv < 0 {
		inv += m
	}
	return inv
}

func Mod(n, base int) int {
	m := n % base
	if m < 0 {
		return m + base
	}
	return m
}
