package formatter

import (
	"testing"

	"github.com/codingingo/leetcode-readme-builder/converter/item"
	"github.com/google/go-cmp/cmp"
)

func TestAlgorithmsFormatter_Format(t *testing.T) {
	type args struct {
		items []item.Item
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"simple",
			args{[]item.Item{{NO: "0430", Question: "Flatten a Multilevel Doububly Linked List", QuestionURL: "https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/", Answer: "0430_flatten_a_multilevel_doubly_linked_list.go", AnswerURL: "https://github.com/codingingo/leetcode/blob/main/0430_flatten_a_multilevel_doubly_linked_list.go", Tags: []string{"linkedlist", "recursion"}}}},
			[]string{"## 标准库\n\n", "| 编号 | 题目 | 题解 |\n| :-: | :-: | :-: |\n", "| 0430 | [Flatten a Multilevel Doububly Linked List](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/) | [0430_flatten_a_multilevel_doubly_linked_list.go](https://github.com/codingingo/leetcode/blob/main/0430_flatten_a_multilevel_doubly_linked_list.go) |\n", "\n---\n\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &AlgorithmsFormatter{}
			got := f.Format(tt.args.items)
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
