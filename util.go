package ossinspector

import (
	"strconv"
	"strings"
)

type ConstraintType int

const (
	GREATER_THAN ConstraintType = iota
	LESSER_THAN
)

func evaluate(val string) (ConstraintType, int) {

	var ctype ConstraintType
	var value int

	if strings.HasPrefix("<", val) {
		s := strings.Split(val, "<")[1]
		num, err := strconv.Atoi(s)
		if err == nil {
			ctype = LESSER_THAN
			value = num

		}
	}

	if strings.HasPrefix(">", val) {
		s := strings.Split(val, ">")[1]
		num, err := strconv.Atoi(s)
		if err == nil {
			ctype = GREATER_THAN
			value = num
		}
	}

	return ctype, value
}
