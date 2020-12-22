package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	con "github.com/codingingo/leetcode-readme-builder/converter"
	"github.com/codingingo/leetcode-readme-builder/converter/filter"
	"github.com/codingingo/leetcode-readme-builder/converter/formatter"
	"github.com/codingingo/leetcode-readme-builder/converter/parser"
)

type Builder interface {
	Build() error
}

func NewBuilder(name string, readPath string, writePath string, converts ...con.Converter) Builder {
	b := &defaultBuilder{}
	b.name = name
	b.readPath = readPath
	b.writePath = writePath
	b.converters = converts
	return b
}

type defaultBuilder struct {
	name       string
	readPath   string
	writePath  string
	converters []con.Converter
}

func (b *defaultBuilder) Build() error {
	// read
	var filenames []string
	err := filepath.Walk(b.readPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	})

	// write
	wt, err := os.OpenFile(b.writePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	// write
	write := bufio.NewWriter(wt)
	write.WriteString(fmt.Sprintf("# %s\n\n", b.name))

	for _, c := range b.converters {
		lines, err := c.Convert(filenames)
		if err != nil {
			return err
		}
		for _, line := range lines {
			write.WriteString(line)
		}
		write.Flush()
	}
	return nil
}

func main() {
	title, readPath, writePath := "Leetcode题解", "./testdata/", "./TEST-README.md"
	algo := con.NewAlgorithmsConverter("标准题库", filter.NewAlgorithmsFilter(), parser.NewAlgorithmsParser(), formatter.NewAlgorithmsFormatter())
	builder := NewBuilder(title, readPath, writePath, algo)
	log.Fatal(builder.Build())
}
