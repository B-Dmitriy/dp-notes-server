package utils

import (
	"fmt"
)

func NilToStr(arg interface{}) string {
	if arg == nil {
		return ""
	} else {
		return fmt.Sprint(arg)
	}
}
