package vanity

import (
	"regexp"
)

type Worker struct {
	Expression regexp.Regexp
}
