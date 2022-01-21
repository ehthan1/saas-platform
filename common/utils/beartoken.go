package utils

import (
	"errors"
	"strings"
)

func ParseBear(tok string) (string, error) {
	if tok == "" {
		return tok, errors.New("token does not exist")
	}
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, errors.New("token invild")
}
