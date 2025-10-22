package other

import (
	"fmt"
	"math/big"
)

/*
Если даны две двоичные строки a и b, верните их сумму в виде двоичной строки.
*/

// runtime: 1ms, beats 24.36%, mem: 4.60 MB

func AddBinary(a string, b string) string {
	num1 := new(big.Int)
	num2 := new(big.Int)

	num1.SetString(a, 2)
	num2.SetString(b, 2)

	sum := new(big.Int).Add(num1, num2)

	return fmt.Sprintf("%b", sum)
}
