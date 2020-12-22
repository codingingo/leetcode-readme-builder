# leetcode-readme-builder

# Just for fun

readme file builder for leetcode

```
		1. 数据来源
		当前目录下*.go文件——文件名+文件中的标记
		1) 读取当前目录下的所有文件名
		2）发送给注册的
		过滤规则：
		Builder
		New(path, targetpath, ....Converter)
		Build()
		read(path) - go converter.convert() -channel- go write(targetpath)
		           - go converter.convert()


		文件名有固定格式，文件内容的头部还有其他信息以tag

		Reader文件名流 -> 过滤规则(分发到自己的builder)->过滤规则2->
						|
					  文件名不同
					   FilterFunc
					   Filter - Apply（filenames []string, ...filterFunc）[]sting
					   Filter - Apply（filenames string, filterFunc）(sting, error)
					   ｜得到文件列表
					   利用自定义的规则，过滤文件名，
					   输入：文件名
					   输出：符合规则文件名
						｜ 按照规则解析成
						Parser - Parse(filenames string[]) struct[]
						输入：符合规则的文件名
						输出：结构体：标题，编号，题目+链接，题解+链接，标签
						 / 填充模版（多样化）
									- template()
									## name
									ln
									table
						 Formatter - format("程序员", struct[]) string[]
						 输入：结构体
						 输出：格式化字符串
						   |	- New("name", filenames[], Filter, Parser, Formatter)
						Converter: - Convert() []string
						names :=filter.Apply(filenames[])
						fm.format(“标题”， parser.Parse(names))

						name => []srings
					Writer写入组件：汇总各个格式化的文件字符串
						   写入README

		文件名过滤
			xxx读取当前目录下*.go文件，拿到文件名
			通过规则函数分发
		标记名过滤
			读取标记

		2. 数据去向，README.md文件，格式化+表格等等

		文件格式化模版：
		默认模版
		定制模版
```