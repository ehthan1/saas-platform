package xerror

import (
	"errors"
	"strings"
)

func ErrorMsg(err error) (errMsg error) {
	return errors.New(strings.TrimLeft(err.Error(), "rpc error: code = Unknown desc ="))
}
