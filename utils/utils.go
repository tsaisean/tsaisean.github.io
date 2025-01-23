package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func GetUnmarshalTypeErrorMsg(err error) (bool, string) {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		message := fmt.Sprintf("Field '%s' has a type mismatch. Expected '%s'.", unmarshalTypeError.Field, unmarshalTypeError.Type)
		return true, message
	}
	return false, ""
}
