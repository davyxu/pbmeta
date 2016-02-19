package pbmeta

import (
	"bytes"
	"strings"
)

type TaggedComment struct {
	Name string
	Data string
}

const (
	stateTagBegin = iota
	stateTagName
	stateMeta
	stateMetaString
)

func parseTaggedComment(src string, commentArray []*TaggedComment) []*TaggedComment {

	lex := newCommentLexer(strings.TrimSpace(src))

	var buff bytes.Buffer
	var state int = stateTagBegin
	var cmt *TaggedComment
	var c byte
	var ok bool

	// 不要新区, 用上次的
	var keep bool

	for {

		if !keep {
			c, ok = lex.Next()

			if !ok {
				break
			}

		} else {
			keep = false
		}

		//fmt.Printf("%d: %d %c\n", lex.pos, state, c)

		switch state {
		case stateTagBegin:
			if c != '[' {
				return commentArray
			}

			cmt = new(TaggedComment)

			buff.Reset()

			state = stateTagName
		case stateTagName:

			if c == ']' {
				state = stateMeta
				cmt.Name = buff.String()
				buff.Reset()
			} else {
				buff.WriteByte(c)
			}

		case stateMeta:

			switch c {
			case '#':
				goto parseEnd
			case '[':
				// 这个已经完成
				cmt.Data = buff.String()
				commentArray = append(commentArray, cmt)

				// 开始探测新的tag开始
				state = stateTagBegin
				keep = true
			// 判断data中有字符串的, 字符串直接输出
			case '"':
				state = stateMetaString
				buff.WriteByte(c)
				// 字符串开始
			default:
				// meta的内容
				buff.WriteByte(c)
			}
		case stateMetaString:

			// 检测字符串结束
			switch c {
			case '"':
				state = stateMeta
				keep = true
			case '\\': // 转义双引号
				c, ok := lex.Peek()
				if !ok {
					goto parseEnd
				}

				if c != '"' {
					buff.WriteByte('\\')
				} else {
					//fmt.Printf("%d': %d %c\n", lex.pos, state, c)

					buff.WriteByte(c)
					lex.Skip()
				}

			default:
				buff.WriteByte(c)
			}
		}

	}

parseEnd:
	if cmt != nil {
		cmt.Data = buff.String()
		commentArray = append(commentArray, cmt)
	}

	return commentArray

}

type commentLexer struct {
	src string
	pos int
}

func (self *commentLexer) Next() (byte, bool) {

	if c, ok := self.Peek(); ok {
		self.pos++
		return c, true
	}

	return 0, false

}

func (self *commentLexer) Peek() (byte, bool) {
	if self.pos >= len(self.src) {
		return 0, false
	}

	c := self.src[self.pos]

	return c, true
}

func (self *commentLexer) Skip() {
	self.pos++
}

func newCommentLexer(src string) *commentLexer {
	return &commentLexer{src: src}
}
