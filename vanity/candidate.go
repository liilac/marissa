package vanity

import (
	"github.com/liilac/marissa/llarp"
	"regexp"
	"strings"
)

type ServiceAddressCandidate llarp.ServiceAddress

func (sac ServiceAddressCandidate) String() string {
	return llarp.ServiceAddress(sac).String()
}

func (sac ServiceAddressCandidate) MatchesRegexp(exp regexp.Regexp) bool {
	return exp.MatchString(sac.String())
}

func (sac ServiceAddressCandidate) MatchesPrefixString(prefix string) bool {
	return strings.ToLower(prefix) == sac.String()[:len(prefix)]
}
