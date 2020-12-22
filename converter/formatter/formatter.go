package formatter

import (
	"fmt"

	"github.com/codingingo/leetcode-readme-builder/converter/item"
)

type Formatter interface {
	Format(items []item.Item) []string
}

func NewAlgorithmsFormatter() Formatter {
	return &AlgorithmsFormatter{}
}

// AlgorithmsFormatter defines algorithm
type AlgorithmsFormatter struct{}

// Format
func (f *AlgorithmsFormatter) Format(items []item.Item) []string {
	rows := []string{}
	header := "## 标准库\n\n"
	table := "| 编号 | 题目 | 题解 |\n| :-: | :-: | :-: |\n"
	rows = append(rows, header, table)
	for _, item := range items {
		rows = append(rows, fmt.Sprintf("| %s | [%s](%s) | [%s](%s) |\n", item.NO, item.Question, item.QuestionURL, item.Answer, item.AnswerURL))
	}
	rows = append(rows, "\n---\n\n")
	return rows
}
