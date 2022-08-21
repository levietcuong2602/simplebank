package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/levietcuong2602/simplebank/utils"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}

	return false
}
