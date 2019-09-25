package parsers

import (
	"bufio"
	"io"

	"github.com/mvdan/xurls"
)

type RegexParser struct{}

func NewRegexParser() *RegexParser {
	return &RegexParser{}
}

func (p *RegexParser) Parse(r io.Reader) ([]string, error) {
	var targets []string
	targetsFilter := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		re := xurls.Relaxed
		for _, target := range re.FindAllString(scanner.Text(), -1) {
			if _, found := targetsFilter[target]; found {
				continue
			}
			targets = append(targets, target)
			targetsFilter[target] = struct{}{}
		}
	}
	return targets, nil
}
