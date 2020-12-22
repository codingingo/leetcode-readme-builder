package parser

import (
	"testing"

	"github.com/codingingo/leetcode-readme-builder/converter/item"
	"github.com/google/go-cmp/cmp"
)

func TestAlgorithmsParser_Parse(t *testing.T) {
	type args struct {
		filenames []string
	}
	tests := []struct {
		name      string
		args      args
		wantItems []item.Item
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			"simple",
			args{[]string{"../../testdata/0430_flatten_a_multilevel_doubly_linked_list.go"}},
			[]item.Item{{NO: "0430", Question: "Flatten a Multilevel Doububly Linked List", QuestionURL: "https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/", Answer: "0430_flatten_a_multilevel_doubly_linked_list.go", AnswerURL: "https://github.com/codingingo/leetcode/blob/main/0430_flatten_a_multilevel_doubly_linked_list.go", Tags: []string{"linkedlist", "recursion"}}},
			false,
		},
		{
			"simple",
			args{[]string{"0430_flatten_a_multilevel_doubly_linked_list"}},
			([]item.Item)(nil),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AlgorithmsParser{}
			gotItems, err := p.Parse(tt.args.filenames)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlgorithmsParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(tt.wantItems, gotItems)
			if diff != "" {
				t.Fatalf(diff)
			}

		})
	}
}
