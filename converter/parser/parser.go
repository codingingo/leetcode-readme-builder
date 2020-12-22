package parser

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/codingingo/leetcode-readme-builder/converter/item"
)

const (
	NO = iota
	QUESTION
	QUESTIONURL
	TAG

	BaseAnswerURL = "https://github.com/codingingo/leetcode/blob/main/"
)

type Parser interface {
	Parse(filenames []string) ([]item.Item, error)
}

func NewAlgorithmsParser() Parser {
	return &AlgorithmsParser{}
}

type AlgorithmsParser struct{}

func (p *AlgorithmsParser) Parse(filenames []string) (items []item.Item, err error) {

	for _, filename := range filenames {

		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		// read header
		n, lines := 0, []string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() && n < 4 {
			lines = append(lines, scanner.Text())
			n++
		}

		item := item.Item{
			NO:          strings.Split(lines[NO], ":")[1],
			Question:    strings.Split(lines[QUESTION], ":")[1],
			QuestionURL: strings.Join(strings.Split(lines[QUESTIONURL], ":")[1:3], ":"),
			Answer:      filepath.Base(file.Name()),
			AnswerURL:   BaseAnswerURL + filepath.Base(file.Name()),
			Tags:        strings.Split(strings.Split(lines[TAG], ":")[1], ","),
		}
		items = append(items, item)
	}

	return
}
