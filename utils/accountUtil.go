package utils

import "github.com/levietcuong2602/simplebank/constants"

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(100, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "VND", "EUR"}
	n := len(currencies)
	return currencies[RandomInt(0, int64(n))]
}

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case constants.EUR, constants.USD, constants.VND:
		return true
	}

	return false
}
