package models

import "gopkg.in/bluesuncorp/validator.v6"

//ValidStruct a given one
func ValidStruct(s interface{}) (bool, validator.ValidationErrors) {
	errorsMap := GetValidator().Struct(s)
	return errorsMap == nil, errorsMap
}

//GetValidator reference
func GetValidator() *validator.Validate {
	config := validator.Config{
		TagName:         "validate",
		ValidationFuncs: validator.BakedInValidators,
	}
	return validator.New(config)
}
