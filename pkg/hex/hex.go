package hex

import "math/big"

const (
	base10 = 10
	base16 = 16
)

func ConvertTo16Bit(value string) string {
	lNumber := new(big.Int)
	lNumber.SetString(value, base10)
	hexString := lNumber.Text(base16)

	return hexString
}
