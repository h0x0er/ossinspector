package ossinspector

import (
	"strconv"
	"strings"
)

type ConstraintType byte

const (
	GREATER_THAN ConstraintType = iota
	LESSER_THAN

	MONTHS_GREATER_THAN
	MONTHS_LESSER_THAN

	YEARS_GREATER_THAN
	YEARS_LESSER_THAN

	DAYS_GREATER_THAN
	DAYS_LESSER_THAN

	NO_CONSTRAINT
)

func evaluate(expr string) (ConstraintType, int) {

	var ctype ConstraintType
	var value int

	// TODO: checking if expr in of age type
	if isAge(expr) {
		value = trim(expr)
		switch expr[len(expr)-1] {
		case 'd':
			ctype = DAYS_GREATER_THAN
			if isLesser(expr) {
				ctype = DAYS_LESSER_THAN
			}

		case 'y':
			ctype = YEARS_GREATER_THAN
			if isLesser(expr) {
				ctype = YEARS_LESSER_THAN
			}

		case 'm':
			ctype = MONTHS_GREATER_THAN
			if isLesser(expr) {
				ctype = MONTHS_LESSER_THAN
			}

		default:

		}

		return ctype, value
	}

	if isLesser(expr) {
		s := strings.Split(expr, "<")[1]
		num, err := strconv.Atoi(s)
		if err == nil {
			ctype = LESSER_THAN
			value = num

		}
	}
	if isGreater(expr) {
		s := strings.Split(expr, ">")[1]
		num, err := strconv.Atoi(s)
		if err == nil {
			ctype = GREATER_THAN
			value = num
		}
	}

	return ctype, value
}

func isAge(expr string) bool {
	last := len(expr) - 1
	return expr[last] == 'd' || expr[last] == 'y' || expr[last] == 'm'
}
func isLesser(expr string) bool {
	return expr[0] == '<'
}
func isGreater(expr string) bool {
	return expr[0] == '>'
}

// trim removes suffix and prefix
// return the middle value, return -1 if failed
func trim(expr string) int {
	tail := expr[1:]
	head := tail[0 : len(tail)-1]
	num, err := strconv.Atoi(head)
	if err == nil {
		return num
	}
	return -1
}
