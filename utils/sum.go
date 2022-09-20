package utils

import "strconv"

// AmountSum returns the payable price and markup
func AmountSum(base int, amount string, isPercent byte) (int, int) {
	if isPercent == '1' {
		m, _ := strconv.Atoi(amount)
		m = base * m / 100
		return base + m, m
	}
	m, _ := strconv.Atoi(amount)
	return base + m, m
}
