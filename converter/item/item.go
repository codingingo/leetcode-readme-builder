package item

// Item defines
// 编号，题目+链接，题解+链接，标签
type Item struct {
	NO          string
	Question    string
	QuestionURL string
	Answer      string
	AnswerURL   string
	Tags        []string
}
