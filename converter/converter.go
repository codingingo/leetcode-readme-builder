package converter

import (
	"github.com/codingingo/leetcode-readme-builder/converter/filter"
	"github.com/codingingo/leetcode-readme-builder/converter/formatter"
	"github.com/codingingo/leetcode-readme-builder/converter/parser"
)

type Converter interface {
	Convert(filenames []string) ([]string, error)
}

func NewAlgorithmsConverter(name string, filter filter.Filter, parser parser.Parser, formatter formatter.Formatter) Converter {
	return &AlgorithmsConverter{name, filter, parser, formatter}
}

type AlgorithmsConverter struct {
	name string
	f    filter.Filter
	p    parser.Parser
	fm   formatter.Formatter
}

func (c *AlgorithmsConverter) Convert(filenames []string) ([]string, error) {
	filtered := c.f.Apply(filenames)
	items, err := c.p.Parse(filtered)
	if err != nil {
		return nil, err
	}
	lines := c.fm.Format(items)
	return lines, nil
}
